# Pattern - Pacade 

## Basic Information 

- 패턴 형식 : 목적 패턴 

## Gof's Description 

한 서브 시스템 내의 인터페이스 집합에 대한 획일화된 하나의 인터페이스를 제공하는 패턴으로, 서브 시스템을 사용하기 쉽도록 상위 수준의 인터페이스를 제공합니다.

## 구조적인 설명 

한 예로, 응용 프로그램이 컴파일러 시스템에 접근할 수 있는 어떤 프로그래밍 환경을 가정했을 때,   
해당 컴파일러의 서브시스템에는 Scanner, Parser, BytecodStream 등의 클래스를 포함합니다.   
이때 응용프로그램의 어느 특정 부분은 컴파일러에 정의된 이 클래스를 직접사용해야할 수 있지만,   
대부분의 응용 프로그램들은 이런 구체적인 내용에 상관없이 파싱이나 코드 생성 단계를 이용하기만 합니다.

## 실생활 예제 

### Java 에서는 

```java 
package lines.loader;

enum ComponentType {
    Heap,
    JavaThreads,
    MethodArea,
    NativeInternalThreads,
    ProgramCounterRegisters
}

public class ClassLoader {

    public static void main(String[] args) {
        JVMComponent.Builder builder = new JVMComponent.Builder();


        builder.Heap(new Heap());
        builder.JavaThreads(new JavaThreads());
        builder.MethodArea(new MethodArea());
        builder.ProgramCounterRegisters(new ProgramCounterRegisters());
        builder.NativeInternalThreads(new NativeInternalThreads());

        ClassLoadSubSystem classLoadSubSystem = new ClassLoadSubSystem(builder.build());

        classLoadSubSystem.ControlVariables();
    }
}

class ClassLoadSubSystem {

    JVMComponent jvmComponent;

    public ClassLoadSubSystem(JVMComponent jvmComponent){
        this.jvmComponent = jvmComponent;
    }

    public void ControlVariables(){

        JVMElement methodArea = jvmComponent.getJVMCompoentOrNull(ComponentType.MethodArea);
        JVMElement heap = jvmComponent.getJVMCompoentOrNull(ComponentType.Heap);
        JVMElement javaThreads = jvmComponent.getJVMCompoentOrNull(ComponentType.JavaThreads);
        JVMElement nativeInternalThreads = jvmComponent.getJVMCompoentOrNull(ComponentType.NativeInternalThreads);
        JVMElement programCounterRegisters = jvmComponent.getJVMCompoentOrNull(ComponentType.ProgramCounterRegisters);


        methodArea.doProcess();
        heap.doProcess();
        javaThreads.doProcess();
        nativeInternalThreads.doProcess();
        programCounterRegisters.doProcess();
    }
}

abstract class JVMElement {
    public abstract void doProcess();
}

class Heap extends JVMElement {

    public void doProcess() {

        forHeapSetting();
    }

    private void forHeapSetting(){

    }
}

class JavaThreads extends JVMElement {
    @Override
    public void doProcess() {

        forJavaThreadSetting();
    }

    private void forJavaThreadSetting(){

    }
}

class MethodArea extends JVMElement {

    public void doProcess() {

        forMethodAreaSetting();
    }

    private void forMethodAreaSetting(){

    }
}

class NativeInternalThreads extends JVMElement {

    public void doProcess() {

        forNativeInternalThreadSetting();
    }

    private void forNativeInternalThreadSetting(){

    }
}

class ProgramCounterRegisters extends JVMElement {

    public void doProcess() {

        forProgramCounterRegisters();
    }

    private void forProgramCounterRegisters(){

    }
}

class JVMComponent {

    private final Heap heap;
    private final JavaThreads javaThreads;
    private final MethodArea methodArea;
    private final NativeInternalThreads nativeInternalThreads;
    private final ProgramCounterRegisters programCounterRegisters;

    static class Builder{

        private Heap heap;
        private JavaThreads javaThreads;
        private MethodArea methodArea;
        private NativeInternalThreads nativeInternalThreads;
        private ProgramCounterRegisters programCounterRegisters;

        public Builder Heap(Heap heap){
            this.heap = heap;
            return this;
        }

        public Builder JavaThreads(JavaThreads javaThreads){
            this.javaThreads = javaThreads;
            return this;
        }

        public Builder MethodArea(MethodArea methodArea){
            this.methodArea = methodArea;
            return this;
        }

        public Builder NativeInternalThreads(NativeInternalThreads nativeInternalThreads){
            this.nativeInternalThreads = nativeInternalThreads;
            return this;
        }

        public Builder ProgramCounterRegisters(ProgramCounterRegisters programCounterRegisters){
            this.programCounterRegisters = programCounterRegisters;
            return this;

        }


        public JVMComponent build(){
            return new JVMComponent(this);
        }

    }

    private JVMComponent(Builder build){
        this.heap = build.heap;
        this.javaThreads = build.javaThreads;
        this.methodArea = build.methodArea;
        this.nativeInternalThreads = build.nativeInternalThreads;
        this.programCounterRegisters = build.programCounterRegisters;
    }

    public JVMElement getJVMCompoentOrNull(ComponentType type){
        switch (type){
            case Heap :
                return heap;
            case JavaThreads:
                return javaThreads;
            case MethodArea:
                return methodArea;
            case NativeInternalThreads:
                return nativeInternalThreads;
            case ProgramCounterRegisters:
                return programCounterRegisters;
        }

        return null;
    }
}
```