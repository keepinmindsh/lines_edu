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

- Pacade 패턴을 통해서 하나의 시스템 내에서 동작하는 복잡한 프로세스를 하나의 기능으로 캡술화해서 굳이 내부 프로세스를 노출할 필요 없이 사용이 가능하다.  

### 추가적으로 예를 들면, 

- MVC 패턴에서의 Controller 내의 다양한 로직이 Controller에서 정의된 하나의 메소드를 통해서 호출되는 것도 Pacade 패턴의 일종이다. 

- Pacade 패턴은
    - 예약 시점에서 예약 확정을 클릭시 이후 내부에서 발생되는 모든 프로세스를 하나의 기능으로 표현할 수 있다
    - 채팅 시스템에서 메세지 생성 함수 호출 이후 발생해야하는 내부 기능은 호출하는 입장에서는 MessageCreate로 이해하는 것이 좋다.


### 예약 시스템의 예약을 Pacade로 구성한다면,

![Pacade Pattern](https://github.com/keepinmindsh/lines_edu/blob/main/assets/pacade_pattern.png)


## 실생활 예제 

### Go 에서는 

```go 
package main

import (
	"compiler/app/builder"
	"compiler/app/compiler"
	"compiler/app/element"
)

func main() {
	compilerBuilder := builder.NewCompilerBuilder()

	compilerBuilder.Parser(element.NewParser())
	compilerBuilder.Scanner(element.NewScanner())
	compilerBuilder.ProgramNode(element.NewNode())

	compiler := compiler.NewCompiler(compilerBuilder)

	compiler.Compile()
}
```

```go 
package domain

type Compiler interface {
	Compile()
}

type Element interface {
	Process()
}

```

```go 
package builder

import "compiler/domain"

type CompilerElement struct {
	scanElement        domain.Element
	parserElement      domain.Element
	programNodeElement domain.Element
}

func NewCompilerBuilder() *CompilerElement {
	return &CompilerElement{}
}

func (c *CompilerElement) Scanner(element domain.Element) *CompilerElement {
	c.scanElement = element
	return c
}

func (c *CompilerElement) Parser(element domain.Element) *CompilerElement {
	c.parserElement = element
	return c
}

func (c *CompilerElement) ProgramNode(element domain.Element) *CompilerElement {
	c.programNodeElement = element
	return c
}

func (c *CompilerElement) GetScanner() domain.Element {
	return c.scanElement
}

func (c *CompilerElement) GetParser() domain.Element {
	return c.parserElement
}

func (c *CompilerElement) GetProgramNode() domain.Element {
	return c.programNodeElement
}
```

```go 
package element

import (
	"compiler/domain"
	"fmt"
)

type Node struct {
}

func (n Node) Process() {
	fmt.Println("Loading node is processing")
}

func NewNode() domain.Element {
	return &Node{}
}
```

```go
package element

import (
	"compiler/domain"
	"fmt"
)

type Parser struct {
}

func (p Parser) Process() {
	fmt.Println("Parser is processing")
}

func NewParser() domain.Element {
	return &Parser{}
}
```

```go 
package element

import (
	"compiler/domain"
	"fmt"
)

type Scanner struct {
}

func (s Scanner) Process() {
	fmt.Println("Scan is processing")
}

func NewScanner() domain.Element {
	return &Scanner{}
}
```

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

### Kotlin 에서는 

- Compiler 실행

```kotlin 
package lines

import lines.builder.CompilerBuilder
import lines.compiler.JavaCompiler
import lines.compiler.Node
import lines.compiler.Parser
import lines.compiler.Scanner
import lines.domain.Compiler

fun main() {
    val compiler : Compiler = JavaCompiler(CompilerBuilder(
        Node(),
        Parser(),
        Scanner(),
    ))

    compiler.Compile()
}
```

- Compiler 빌더 ( Pacade 구현 용도 )

```kotlin
package lines.builder

import lines.domain.Element

data class CompilerBuilder (
    var parser: Element,
    var scanner: Element,
    var node: Element
)
```

- 도메인 정의 

```kotlin
package lines.domain

interface Compiler {
    fun Compile()
}
```

```kotlin
package lines.domain

interface Element {
    fun Process()
}
```

- Compiler ( Pacade ) 구현체 

```kotlin 
package lines.compiler

import lines.builder.CompilerBuilder

class JavaCompiler(var compiler: CompilerBuilder) : lines.domain.Compiler {
    override fun Compile() {
        compiler.node.Process()

        compiler.parser.Process()

        compiler.scanner.Process()
    }
}
```

- Compiler 요소 구현제 

```kotlin
package lines.compiler

import lines.domain.Element

class Node : Element {
    override fun Process() {
        println("Node is processing")
    }
}
```

```kotlin 
package lines.compiler

import lines.domain.Element

class Parser : Element {
    override fun Process() {
        println("Parser is processing")
    }
}
```

```kotlin
package lines.compiler

import lines.domain.Element

class Scanner : Element {
    override fun Process() {
        println("Scanner is processing")
    }
}
```

## 활용성 

- 복잡한 서브 시스템에 대한 단순한 인터페이스 제공이 필요할 때, 시스템 범위가 확장되면, 또한 구체적으로 설계되면 서브 시스템은 계속 복잡해집니다. 또한 패턴을 적용하면 확장성을 고려하여 설계되기 때문에, 작은 클래스가 만들어지게 됩니다. 이런 과정은 서브 시스템을 재사용 가능한 것으로 만들어 주고, 재정의 할 수 있는 단위가 되도록 해 주기도 하지만, 실제 이런 상세한 재설계나 정제의 내용까지 파악할 필요가 없는 개발자들에게 복잡해진 각각의 클래스들을 다 이해하면서 서브시스템을 다 사용하기란 어려운 일입니다. 이럴 때 퍼사드 패턴은 서브시스템에 대한 단순하면서도 기본적인 인터페이스를 제공함으로써 대부분의 개발자들에게 적합한 클래스 형태를 제공합니다.

- 추상 개념에 대한 구현 클래스와 사용자 사이에 너무 많은 종속성이 존재할 때,

- 서브 시스템을 계층화 시킬 때, 퍼사드 패턴을 사용하여 각 서브 시스템의 계층에 대한 접근점을 제공합니다.

- 서브 시스템의 구성 요소를 보호할 수 있으며, 서브 시스템과 사용자 코드 간의 결합도를 약하게 만들 수 있다.
