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