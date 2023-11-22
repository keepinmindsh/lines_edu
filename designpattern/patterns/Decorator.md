# Pattern - Decorator 

## Basic Information 

- 패턴형식 : 목적패턴 

## Gof's Description 

객체에 동적으로 새로운 책임을 추가할 수 있게 합니다. 기능을 추가하기 위해서 서브클래스를 생성하는 것보다 더욱 융통성있는 방법을 제공합니다.

## 구조적인 설명 

아래의 예시는 TextView로 구현된 객체에 대해서 TextView 가 색상을 추가해야는 요건이 생겼을 경우, 객체를 어떻게 확장할 것인가에 대해서, 

- 서브클래싱을 통한 구현 방식 
- 의존성 주입을 통한 구현 방식 

2가지 방식 중에 더 좋은 방법은 **의존성 주입**을 통한 방식으로 기존의 객체에 영향도를 미치지 않으면서 신규 구성된 객체에서 신규 기능을 구현하는 방식을 추천한다.   

해당 부분의 중요한 부분은 신규 기능도 기존의 기능과 동일한 기능 내에서 세부 설정이 조정되는 범위에서 사용이 가능하다.   

기본적으로 인터페이스를 통한 추상화를 정의할 때 데코레이터의 대상객체와 데코레이터 객체는 동일한 인터페이스를 구현하는 방식을 통해서 기존의 서비스의 기능을   
확장해 나가는 방식이기 때문이다. 

### 잘못된 예제 

![Decorator Bad Sample](https://github.com/keepinmindsh/lines_edu/blob/main/assets/decorator_badsample.png)

### 잘 구성된 예제 

![Decorator Good Sample](https://github.com/keepinmindsh/lines_edu/blob/main/assets/decorator_sample01.png)

## 필요한 경우 

- 동적으로 또한 투명하게, 다시 말해 다른 객체에 영향을 주지 않고 개개의 객체에 새로운 책임을 추가하기 위해서 사용합니다.

- 제거될 수 있는 책임에 대해서 사용됩니다.

- 실제 상속으로 서브클래스를 계속 만드는 방법이 실질적이지 못할 때 사용합니다. 너무 많은 수의 독립된 확장이 가능할 때 모든 조합을 지원하기 위해 이를 상속으로 해결하면 클래스 수가 폭발적으로 많아지게 됩니다. 아니면, 클래스 정의가 숨겨지든가, 그렇지 않더라도 서브클래싱을 할 수 없게 됩니다.

- 장식자 패턴은 객체에 새로운 행동을 추가할 수 있는 가장 효과적인 방법입니다. 장식자를 사용하면 장식자를 객체와 연결하거나 분리하는 작업을 통해 새로운 책임을 추가하거나 삭제하는 일이 런타임에 가능해집니다.


## 실생활 예제 


### Java 에서는 


```java 
package DesignPattern.gof_decorator;

public class Barrack {

    public static void main(String[] args) {
        Upgrade marinUpgrade = new MarineBasicUprade();

        marinUpgrade.upgrade();

        System.out.println("------------");

        Upgrade steampackUpgrade = new SteamPack(marinUpgrade);

        steampackUpgrade.upgrade();

        System.out.println("------------");

        Upgrade longRangeUpgrade = new LongRange(steampackUpgrade);

        longRangeUpgrade.upgrade();

        System.out.println("------------");

        Upgrade speedUpgrade = new Speed(longRangeUpgrade);

        speedUpgrade.upgrade();

        System.out.println("------------");

    }
}
```

```java 
package DesignPattern.gof_decorator;

public abstract class Upgrade {
    public abstract void upgrade();
}
                                    
package DesignPattern.gof_decorator;

public class MarineBasicUprade extends Upgrade {

    @Override
    public void upgrade() {
        System.out.println("공격력이 +1 증가하였습니다.");
    }
}

package DesignPattern.gof_decorator;

public abstract class UpgradeDecorator extends Upgrade {

    protected Upgrade upgrade;

    public UpgradeDecorator(Upgrade upgrade){
        this.upgrade = upgrade;
    }

    public void upgrade(){
        upgrade.upgrade();
    }
}

package DesignPattern.gof_decorator;

public class SteamPack extends UpgradeDecorator {
    public SteamPack(Upgrade upgrade) {
        super(upgrade);
    }

    @Override
    public void upgrade() {
        super.upgrade();

        applySteamPack();
    }

    public void applySteamPack(){
        System.out.println("스팀팩이 적용되었습니다. ");
    }
}

package DesignPattern.gof_decorator;

public class Speed extends UpgradeDecorator {

    public Speed(Upgrade upgrade){
        super(upgrade);
    }

    public void upgrade() {
        super.upgrade();

        applySteamUpgrade();
    }

    public void applySteamUpgrade(){
        System.out.println("속도가 10 만큼 증가하였습니다.");
    }
}

package DesignPattern.gof_decorator;

public class LongRange extends UpgradeDecorator {
    public LongRange(Upgrade upgrade) {
        super(upgrade);
    }

    @Override
    public void upgrade() {
        super.upgrade();

        applyLongRange();
    }

    public void applyLongRange(){
        System.out.println("공격 사거리가 10 증가하였습니다.");
    }
}
```