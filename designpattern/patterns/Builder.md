# Pattern - Builder 

## Basic Information

- 패턴형식 : 생성패턴 

## Gof's Description

### 의도 

생성자의 인자가 많을 때는 Builder 패턴을 고려 해야 합니다.  
빌더 패턴은 인자가 많은 생성자나 정적 팩터리가 필요한 클래스를 설계할 때, 특히 대부분의 인자가 선택적 인자인 상황에 유용 합니다.  

### 풀이

Builder 를 사용해야 하는 주요한 경우는 내가 개발 하는 비즈니스 로직 내의 인자, 파라미터가 명시적으로 정해져 있지 않고, 추후의 변경 가능성을 고려 했을 때, 
함수 파라미터의 인자를 늘리는 방식이 아닌 여러 인자를 담고 있는 인자 하나를 **전달** 하여 추후의 변경 영향도에 대응하는 것 입니다. 

### 장점 

- 제품에 대한 내부 표현을 다양하게 표현할 수 있습니다.
- 생성과 표현에 필요한 코드를 분리합니다. 빌더 패턴을 사용하면 복합 객체를 생성하고 복합 객체의 내부 표현 방법을 별도의 모듈로 정의할 수 있습니다.
- 복합 객체를 생성하는 절차를 좀더 세밀하게 나눌 수 있습니다.


## 실생활의 사용에 대한 예시 

### Go 에서는 


### Java 에서는 

```java 
package DesignPattern.gof_builder.sample01;
    
public class DropShip {

private  int dropshipCapacity = 8;

private  int SeigeTankSize = 0;
private  int MarineSize = 0;
private  int MedicSize = 0;
private  int FireBatSize = 0;
private  int VultureSize = 0;
private  int GhostSize = 0;
private  int GoliathSize = 0;

private  boolean SeigeTankOnBoard = false;
private  boolean MarineOnBoard = false;
private  boolean MedicOnBoard = false;
private  boolean FireBatOnBoard = false;
private  boolean VultureOnBoard = false;
private  boolean GhostOnBoard = false;
private  boolean GoliathOnBoard = false;

  public static class Builder {

      // 필수 값
      private final int dropshipCapacity;

      private int SeigeTankSize = 4;
      private boolean SeigeTankOnBoard = false;
      private int MarineSize = 1;
      private boolean MarineOnBoard = false;
      private int MedicSize = 1;
      private boolean MedicOnBoard = false;
      private int FireBatSize = 1;
      private boolean FireBatOnBoard = false;
      private int VultureSize = 2;
      private boolean VultureOnBoard = false;
      private int GhostSize = 1;
      private boolean GhostOnBoard = false;
      private int GoliathSize = 2;
      private boolean GoliathOnBoard = false;

      public Builder(int dropshipCapacity){
          this.dropshipCapacity = dropshipCapacity;
      }

      public Builder SeigeTank(int SeigeTankCount){
          SeigeTankSize = SeigeTankSize * SeigeTankCount;
          SeigeTankOnBoard = true;
          return this;
      }

      public Builder Marine(int MarineCount){
          MarineSize = MarineSize * MarineCount;
          MarineOnBoard = true;
          return this;
      }

      public Builder Medic(int MedicCount){
          MedicSize = MedicSize * MedicCount;
          MedicOnBoard = true;
          return this;
      }

      public Builder FireBat(int FireBatCount){
          FireBatSize = FireBatSize * FireBatCount;
          FireBatOnBoard = true;
          return this;
      }

      public Builder Vulture(int VultureCount){
          VultureSize = VultureSize * VultureCount;
          VultureOnBoard = true;
          return this;
      }

      public Builder Ghost(int GhostCount){
          GhostSize = GhostSize * GhostCount;
          GhostOnBoard = true;
          return this;
      }

      public Builder Goliath(int GoliathCount){
          GoliathSize = GoliathSize * GoliathCount;
          GoliathOnBoard = true;
          return this;
      }

      public DropShip onBoard(){
          return new DropShip(this);
      }
  }

  private DropShip(Builder builder){
      dropshipCapacity    = builder.dropshipCapacity;
      SeigeTankSize       = builder.SeigeTankSize;
      MarineSize          = builder.MarineSize;
      MedicSize           = builder.MedicSize;
      FireBatSize         = builder.FireBatSize;
      VultureSize         = builder.VultureSize;
      GhostSize           = builder.GhostSize;
      GoliathSize         = builder.GoliathSize;

      SeigeTankOnBoard       = builder.SeigeTankOnBoard;
      MarineOnBoard          = builder.MarineOnBoard;
      MedicOnBoard           = builder.MedicOnBoard;
      FireBatOnBoard         = builder.FireBatOnBoard;
      VultureOnBoard         = builder.VultureOnBoard;
      GhostOnBoard           = builder.GhostOnBoard;
      GoliathOnBoard         = builder.GoliathOnBoard;
  }

  public String checkOnBoardMember(){
      String onBoardList = "";

      if(SeigeTankOnBoard) onBoardList += " 시즈탱크 탑승 : " + SeigeTankSize + "\r\n";
      if(MarineOnBoard) onBoardList += " 마린 탑승 : " + MarineSize + "\r\n";
      if(MedicOnBoard) onBoardList += " 메딕 탑승 : " + MedicSize + "\r\n";
      if(FireBatOnBoard) onBoardList += " 파이어벳 탑승 : " + FireBatSize + "\r\n";
      if(VultureOnBoard) onBoardList += " 벌처 탑승 : " + VultureSize + "\r\n";
      if(GhostOnBoard) onBoardList += " 고스트 탑승 : " + GhostSize + "\r\n";
      if(GoliathOnBoard) onBoardList += " 골리앗 탑승 : " + GoliathSize + "\r\n";

      return onBoardList;
  }
}   
```

```java  
package DesignPattern.gof_builder.sample01;

public class OnBoardMain {
    public static void main(String[] args) {
        DropShip.Builder onBoardOrder = new DropShip.Builder(8);

        onBoardOrder.FireBat(1);
        onBoardOrder.Ghost(1);
        onBoardOrder.SeigeTank(1);
        onBoardOrder.Goliath(1);

        DropShip dropShip = onBoardOrder.onBoard();

        System.out.printf(dropShip.checkOnBoardMember());
    }
}     
```

#### Java에서 Lombok Plugin을 이용할 경우 

```java 
package DesignPattern.gof_builder.sample02;
         
import lombok.Builder;
import lombok.Getter;
import lombok.ToString;

@Getter
@ToString
public class DropShip {
    private int dropshipCapacity;
    private int SeigeTankSize;
    private int MarineSize;
    private int MedicSize;
    private int FireBatSize;
    private int VultureSize;
    private int GhostSize;
    private int GoliathSize;

    @Builder
    public DropShip(int dropshipCapacity, int SeigeTankSize, int MarineSize, int MedicSize, int FireBatSize, int VultureSize, int GhostSize, int GoliathSize){
        dropshipCapacity    = dropshipCapacity;
        SeigeTankSize       = SeigeTankSize;
        MarineSize          = MarineSize;
        MedicSize           = MedicSize;
        FireBatSize         = FireBatSize;
        VultureSize         = VultureSize;
        GhostSize           = GhostSize;
        GoliathSize         = GoliathSize;
    }
}  
```

### Kotlin 에서는 

```kotlin 
package bong.lines.patterns.builder.datas

data class Group(
    var name: String,
    var company: Company,
    var members: List<Member>
)

data class Company(
    var name : String = ""
)

data class Member (
    var name : String,
    val alias : String,
    var year: Int
)
```

```kotlin 
package bong.lines.patterns.builder

import bong.lines.patterns.builder.datas.Company
import bong.lines.patterns.builder.datas.Group
import bong.lines.patterns.builder.datas.Member


class MemberBuilder {
    private var name: String = ""
    private var alias: String = ""
    private var year: Int = 0

    fun name(lambda: () -> String) {
        name = lambda()
    }

    fun alias(lambda: () -> String) {
        alias = lambda()
    }

    fun year(lambda: () -> Int) {
        year = lambda()
    }

    fun build() = Member(name, alias, year)
}


class MemberListBuilder {
    private val employeeList = mutableListOf<Member>()

    fun member(lambda: MemberBuilder.() -> Unit) =
        employeeList.add(MemberBuilder().apply(lambda).build())

    fun build() = employeeList
}

class CompanyBuilder {
    private var name = ""

    fun name(lambda: () -> String) {
        this.name = lambda()
    }

    fun build() = Company(name)
}

class GroupBuilder {

    private var name = ""

    private var company = Company("")

    private val employees = mutableListOf<Member>()

    fun name(lambda: () -> String) {
        name = lambda()
    }

    fun company(lambda: CompanyBuilder.() -> Unit) {
        company = CompanyBuilder().apply(lambda).build()
    }

    fun members(lambda: MemberListBuilder.() -> Unit) =
        employees.addAll(MemberListBuilder().apply(lambda).build())

    fun build() = Group(name, company, employees)
}


fun group(lambda: GroupBuilder.() -> Unit): Group {
    return GroupBuilder().apply(lambda).build()
}
```

```kotlin 
package bong.lines.patterns.builder

val redVelvet =
    group {
        name { "레드벨벳" }
        company {
            name { "SM Entertainment" }
        }
        members {
            member {
                name { "슬기" }
                alias { "곰슬기" }
                year { 1994 }
            }
            member {
                name { "아이린" }
                alias { "얼굴 천재" }
                year { 1991 }
            }
            member {
                name { "웬디" }
                alias { "천사" }
                year { 1994 }
            }
        }
    }

fun main() {
    print("${redVelvet}")
} 
```