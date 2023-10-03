
---
title: "React JS"   
excerpt: "SPA 개발 및 CBD 기반의 개발을 위한 웹 프레임워크"  

categories:  
- Script  
  tags:  
- Script  
  classes: wide  
  last_modified_at: 2021-06-29T22:49:00-05:00  
---

> 다스림의 표준 - 저절로 다스려지게 하라.

***

# GetStarted 

- [Learn ReactJS](https://react.dev/learn)

# create-react-app를 이용하여 react-project 생성하기

```
# create-react-app 을 이용한 작업 환경 구성하기   
# npm {명령어} -{옵션} {Package}           
# install Package 설치하기   
# -g 전역에 설치하기                    
> npm install -g create-react-app  


> create-react-app hello-world      
> cd hello-world      
> npm start   
```

# 역방향 데이터 흐름 추가하기 

```javascript
const App = () => {
    const [filtered, setFilteredText] = useState("")
    
    const onEnteredFilteredText = (event) => {
        setFilteredText(event.target.value)
    }

    return (
    <div className="App">
      <header className="App-header">
        <Search filtered={filtered} onEnteredFilteredText={onEnteredFilteredText} />
        <Table filtered={filtered} />
      </header>
    </div>
    );
}

export default App;
```

```javascript
const Search = (props) => {
    return <div><input value={props.filtered} onChange={props.onEnteredFilteredText} ></input></div>
}

export default Search
```

```javascript
const Table = (props) => {
    return (
        <div>
            {
                props.filtered ? <div> filtered text </div> : <></>
            }
            <div>table1</div>
            <div>table2</div>
            <div>table3</div>
            <div>table4</div>
            <div>table5</div>
        </div>
    )
}

export default Table
```

# 제어 흐름 제어하기 

displayTables 함수 내에서 isChecked 상태 변수에 의해서 Component가 Rendering될 때 if/else 문에서 component를 개별로 정의하여 반환하는 코드를 작성할 수 있습니다. 

```javascript
const Table = (props) => {
    const displayTables = () => {
        let displayTable;
        if(props.isChecked) {
            displayTable = <>
                <div>table1</div>
                <div>table2</div>
                <div>table3</div>
                <div>table4</div>
                <div>table5</div>
            </>
        }else {
            displayTable = <></>
        }

        return displayTable
    }

    return (
        <div>
            {displayTables()}
        </div>
    )
}

export default Table
```

# props 와 state를 나누어 사용하도록 한 이유

먼저 개발자들에게 명확한 관념 모델(static mental model)을 제공합니다.
관념 모델은 무엇이 어떻게 동작하는지 이해할 때 진행되는 일련의 사고 프로세스를 의미합니다.
즉, 논리적으로 이치에 맞는 사고 모델을 제공한다는 것이다.


# create-react-app 을 이용하지 않고 빌드없이 바로 이용하기

```html

<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta http-equiv="X-UA-Compatible" content="ie=edge">
  <title>Document</title>
</head>
<body>
    <div id="root"></div>
    <script src="https://unpkg.com/react@15/dist/react.js"></script>
    <script src="https://unpkg.com/react-dom@15/dist/react-dom.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/babel-core/5.8.34/browser.js"></script>
    <script src="./js/app.js" type="text/babel"></script> 
</body>
</html>

```

```javascript

// js/app.js 
ReactDOM.render(
  <h1>Hello, world!</h1>,
  document.getElementById('root')
);

```

# React Renering Process

### ReactDOM의 역할
React.js는 자바스크립트 내에 DOM Tree와 같은 구조체를 VIRTUAL DOM으로 갖고 있습니다. 다시 그릴 때는 그 구조체의 전후 상태를 비교하여 변경이 필요한 최소한의 요소만 실제 DOM에 반영합니다. 따라서 무작위로 다시 그려도 변경에 필요한 최소한의 DOM만 갱신되기 때문에 빠르게 처리할 수 있습니다.

![](https://keepinmindsh.github.io/lines/assets/img/reactRendering.png){: .align-center}

Virtual DOM 은 DOM 차원에서의 더블 버퍼링이랑 다름이 없다.  
변화가 일어나면 그것을 오프라인 DOM 트리에 적용시키는데,  
이 DOM 트리는 렌더링도 되지 않기때문에 연산 비용이 적다.  
연산이 끝나고나면 그 최종적인 변화를 실제 DOM 에 던져주는 것이다.  
그러면, 레이아웃 계산과 리렌더링의 규모는 커지겠지만, 딱 한번만 연산이 일어나게 된다.  
바로 이렇게, 하나로 묶어서 적용시키는것이, 연산의 횟수를 줄일 수 있는 방법이다.


React가 DOM보다 빠른 것이 아니라, 유지보수 가능한 어플리케이션을 만드는 것을 도와주고 대부분의 경우에 '충분히 빠르다.'  
실제로 최적화 작업을 직접 했을 때가 React를 사용했을 때 보다 빠르지만 이를 자동화하여 처리해주는 React가 유지보수 및 생산성을 고려하였을 때 배가 되는 것  
React 를 사용한다고 해도 최적화 작업이 제대로 이뤄지지 않으면 오히려 속도가 더 느릴수도 있다.

# Props

Props는 프로퍼티 용어의 약어 입니다. props는 컴포넌트에서 전달해서 보관하길 원하는 데이터 입니다. 즉, 컴포넌트 내에서 보관되면, 이 데이터는 수정되지 않고 보존되어야 하는 법칙이 성립된다. 만약 props의 값을 변경하고자 할 때는 컴포넌트 내부가 아닌, 부모 컴포넌트에서 이에 대한 부분이 변경되어야 합니다.

- props는 읽기 전용입니다.
- 모든 React 컴포넌트는 자신의 props를 다룰 때 반드시 순수 함수 처럼 동작해야 합니다.
- props의 이름은 사용될 context가 아닌 컴포넌트 자체의 관점에서 짓는 것을 권장합니다.

props는 React에서는 사용자가 컴포넌트에서 전달해서 보관하길 원하는 데잍어입니다. 즉, 컴포넌트 내에서 데이터가 보관되면. 이 데이터는 수정되지 않고 보존되어야 하는 법칙이 성립됩니다.

```javascript

function Welcome(props) {
  return <h1>Hello, {props.name}</h1>;
}

const element = <Welcome name="Sara" />;
ReactDOM.render(
  element,
  document.getElementById('root')
);

```

```javascript

// App.js
import React from 'react';
import Hello from './Hello';

function App() {
  return (
    <Hello name="react" />
  );
}

export default App;

// Hello.js
import React from 'react';

function Hello(props) {
  return <div>안녕하세요 {props.name}</div>
}

export default Hello;

```

```javascript

// App.js
import React from 'react';
import Hello from './Hello';

function App() {
  return (
    <Hello name="react" color="red"/>
  );
}

export default App;

// Hello.js
import React from 'react';

function Hello({ color, name }) {
  return <div style={{ color }}>안녕하세요 {name}</div>
}

export default Hello;

```

### defaultProps 로 기본값 설정

```javascript 

// Hello.js
import React from 'react';

function Hello({ color, name }) {
  return <div style={{ color }}>안녕하세요 {name}</div>
}

Hello.defaultProps = {
  name: '이름없음'
}

export default Hello;

// App.js
import React from 'react';
import Hello from './Hello';

function App() {
  return (
    <>
      <Hello name="react" color="red"/>
      <Hello color="pink"/>
    </>
  );
}

export default App;

```


# State

동적인 데이터를 사용해야할 경우, state를 사용해야합니다. 아래의 새로운 컴포넌트를 만들어 보겠습니다.

```javascript 

import React, { Component } from 'react';

class CounterHandler extends Component {
  state = {
    count: 0
  }

  addCount = () => {
    this.setState({
      count: this.state.count + 1
    });
  }

  reduceCount = () => {
    this.setState({
      count: this.state.count - 1
    });
  }

  render() {
    return (
      <div>
        <h1>카운터</h1>
        <div>값: {this.state.count}</div>
        <button onClick={this.addCount}>+</button>
        <button onClick={this.reduceCount}>-</button>
      </div>
    );
  }
}

export default CounterHandler;


```

```javascript 

import React, { Component } from 'react';
import CounterHandler from './CounterHandler';

ReactDOM.render(
    <CounterHandler />
, document.getElementById('root'));


```
React 컴포넌트는 컴포넌트의 상태를 저장할 수 있습니다.  
props와의 차이점이라고 하면, state는 컴포넌트의 내부에 존재하고 있기 때문에, 상태 값 변경이 가능하다는 것입니다.  
this.setState() 메소드를 통해서 stat 값을 변경해줄 수 있습니다.


기본적으로 아래와 같이 Contructor에 state 변수를 선언할 때 중첩 문은 동작하지 않습니다.

```javascript

// 아래와 같이 state 에서 중첩문이 적용되지 않음 - value가 undefined로 인식됨                             
class App extends Component {
  construct(props){
    super(props);
    this.state = {
      result : {
        value : 10
      }
    }
  }
}

```

React는 성능을 위해 여러 setState() 호출을 단일 업데이트로 한꺼번에 처리할 수 있습니다.  
this.props 와 this.state가 비동기적으로 업데이트 될 수 있기 때문에 다음 state를 계산할 때 해당 값에 의존해서는 안됩니다.

### React Hook을 활용한 state 사용 방식

```javascript

// Correct
const [ value, setValue ] = useState(0);

setValue((state, props) => ({
  counter: state.counter + props.increment
}))

```

State는 다양한 독립적인 변수를 포함할 수 있습니다. 위의 코드에서 setValue를 이용해서 독립적으로 변수를 각각 호출할 수 있습니다.

```javascript

// Correct
const [ value, setValue ] = useState({ posts: [], comments: [] });

fetchPosts().then(response => {
  setValue({
    posts: response.posts
  });
});

fetchComments().then(response => {
  setValue({
    comments: response.comments
  });
});

```

state가 소유하고 설정한 컴포넌트 이외에는 어떠한 컴포넌트에서 접근할 수 없습니다.
일반적으로 이를 “하향식(top-down)” 또는 “단방향식” 데이터 흐름이라고 합니다. 모든 state는 항상 특정한 컴포넌트가 소유하고 있으며 그 state로부터 파생된 UI 또는 데이터는 오직 트리구조에서 자신의 “아래”에 있는 컴포넌트에만 영향을 미칩니다.

```javascript 

import React, { useState } from 'react';

const FormattedData = props => {
    return <h2>It is {props.date.toLocaleTimeString()}. </h2>
}

const SampleScreen01 = props => {

    const [ date , setDate] = useState(new Date());

    setInterval(() => {
        setDate(new Date());
      }, 1000)

    // 모든 컴포넌트가 완전히 독립적이라는 것을 보여주기 위해 App 렌더링하는 세 개의 <FormattedData>을 만들었습니다.
    let SampleScreen01Layout = (
        <div className="row" >
          <FormattedData date={date} />
          <hr />
          <FormattedData date={date} />
          <hr />
          <FormattedData date={date} />
          <hr />
        </div>
    );

    return SampleScreen01Layout;
}

export default SampleScreen01;

```

# React Ref

Ref의 바람직한 사용 사례
- 포커스, 텍스트 선택 영역, 혹은 미디어의 재생을 관리할 때,
- 애니메이션을 직접적으로 실행시킬 때,
- 서드파티 DOM 라이브러리를 React와 같이 사용할 때,

Ref 생성하기

Ref는 React.createRef를 통해 생성되고, ref 어트리뷰트를 통해 React 엘리먼트에 부착됩니다. 보통, 컴포넌트의 인스턴스가 생성될 대 Ref를 프로퍼티로 추가하고, 그럼으로서 컴포넌트의 인스턴스 어느 곳에서도 Ref가 접근할 수 있게 합니다.

```javascript


class MyComponent extends React.Component {
  constructor(props) {
    super(props);
    this.myRef = React.createRef();
  }
  render() {
    return <div ref={this.myRef} />;
  }
}

```

### Forwarding refs in higher-order components

고차 함수에서 ref를 전달하고 싶을 때, props는 전달이 가능하지만 ref는 props가 아니기 때문에 일반적인 방법으로는 ref를 전달할 수 없습니다. 하지만 우리는 forwardedRef를 통해서 고차 컴포넌트에 해당 ref 값을 전달 할 수 있게 됩니다.

```javascript

import FancyButton from './FancyButton';

const ref = React.createRef();

// The FancyButton component we imported is the LogProps HOC.
// Even though the rendered output will be the same,
// Our ref will point to LogProps instead of the inner FancyButton component!
// This means we can't call e.g. ref.current.focus()
<FancyButton
  label="Click Me"
  handleClick={handleClick}
  ref={ref}
/>; 


function logProps(Component) {
  class LogProps extends React.Component {
    componentDidUpdate(prevProps) {
      console.log('old props:', prevProps);
      console.log('new props:', this.props);
    }

    render() {
      const {forwardedRef, ...rest} = this.props;

      // Assign the custom prop "forwardedRef" as a ref
      return <Component ref={forwardedRef} {...rest} />;
    }
  }

  // Note the second param "ref" provided by React.forwardRef.
  // We can pass it along to LogProps as a regular prop, e.g. "forwardedRef"
  // And it can then be attached to the Component.
  return React.forwardRef((props, ref) => {
    return <LogProps {...props} forwardedRef={ref} />;
  });
}

```


# JSX

JSX : Javascript + XML 을 합쳐서 JSX라고 말한다.

```javascript 

const element = <h1>Hello, world!</h1>;

ReactDOM.render(
  element,
  document.getElementById('root')
);

```

위의 코드에서 볼 수 있듯이 마치 html 의 요소를 정의한 것처럼 보이지만 이는 모두 JSX에 선언된 요소들이라고 보는게 옳다.

```javascript 

function getHelloWorldByYourName(user){
  return user.myName;
};

const user = {
    myName : "홍길동"
};

const element = (
  <h1><strong>My Name is {getHelloWorldByYourName(user)}</strong></h1>
);

ReactDOM.render(
  element,
  document.getElementById('root')
);

```

### JSX에서의 표현식

```javascript 

const name = 'Jeong Seung Hwa';
const element = <h1>Hello, {name}</h1>;

ReactDOM.render(
  element,
  document.getElementById('root')
);

```

### JSX에서의 요소 바인딩

```javascript 

function formatName(user) {
  return user.firstName + ' ' + user.lastName;
}

const user = {
  firstName: 'Jeong',
  lastName: 'Seung Hwa'
};

const element = (
  <h1>
    Hello, {formatName(user)}!
  </h1>
);

ReactDOM.render(
  element,
  document.getElementById('root')
);

```

### 속성에 대한 정의

```javascript

// 따옴표를 이용햇 문자열 리터럴을 정의할 수 있습니다. 
const element = <div tabIndex="0"></div>;

// 중괄호를 사용하여 어트리뷰트에 JavaScript 표현식을 삽입할 수도 있습니다. 
const element = <img src={user.avatarUrl}></img>;

```

### JSX로의 자식 정의

```javascript

// tag 안의 자식 요소를 포함할 수 있습니다. 
const element = (
  <div>
    <h1>Hello!</h1>
    <h2>Good to see you here.</h2>
  </div>
);      


```

- JSX는 주입 공격을 방지합니다.
- 기본적으로 React DOM은 JSX에 삽입된 모든 값을 렌더링 하기 전에 이스케이프하므로, 애플리케이션에서 명시적으로 작성되지 않은 내용은 주입되지 않습니다. 모든 항목은 렌더링 되기 전에 문자열로 변환됩니다.
- JSX는 객체를 표현합니다.

```javascript 

// 아래의 두개의 코드는 모두 createElment를 호출하는 방식이지만 표현식이 다릅니다. 
const element = (
  <h1 className="greeting">
    Hello, world!
  </h1>
);

const element = React.createElement(
  'h1',
  {className: 'greeting'},
  'Hello, world!'
);

```

### JSX 속성 요소

```javascript 

accept acceptCharset accessKey action allowFullScreen allowTransparency alt
async autoComplete autoFocus autoPlay capture cellPadding cellSpacing challenge
charSet checked classID className colSpan cols content contentEditable
contextMenu controls coords crossOrigin data dateTime default defer dir
disabled download draggable encType form formAction formEncType formMethod
formNoValidate formTarget frameBorder headers height hidden high href hrefLang
htmlFor httpEquiv icon id inputMode integrity is keyParams keyType kind label
lang list loop low manifest marginHeight marginWidth max maxLength media
mediaGroup method min minLength multiple muted name noValidate nonce open
optimum pattern placeholder poster preload radioGroup readOnly rel required
reversed role rowSpan rows sandbox scope scoped scrolling seamless selected
shape size sizes span spellCheck src srcDoc srcLang srcSet start step style
summary tabIndex target title type useMap value width wmode wrap


```

# Element

```javascript

const element = <h1> Hello, World </h1>;
// 브라우저의 DOM 엘리먼트와 달리 React 엘리먼트는 일반 객체이며 쉽게 생성할 수 있습니다. 
// React DOM 은 React 엘리먼트와 일치하도록 DOM을 업데이트합니다. 

```

### DOM에 엘리먼트 렌더링하기

```javascript

// 이 안에 들어가는 모든 엘리먼트를 React DOM에서 관리하기 때문에 이것을 "루트(root)" DOM 노드라고 부릅니다. 
<div id="root" > <div>

// React 엘리먼트를 Root 엘리먼트에 로딩하려면 아래와 같이 사용합니다. 
const element = <h1>Hello, world</h1>;
ReactDOM.render(element, document.getElementById('root'));

```

### 렌더링 된 엘리먼트 업데이트 하기

```javascript

// ReactDOM 의 render에 의해서 UI 요소를 갱신하는 방식                             
function tick() {
  const element = (
    <div>
      <h1>Hello, world!</h1>
      <h2>It is {new Date().toLocaleTimeString()}.</h2>
    </div>
  );
  ReactDOM.render(element, document.getElementById('root'));
}

setInterval(tick, 1000);

```

### Components

컴포넌트 기반 소프트웨어 공학 ( Component-based software engineering)

CBD는 개발 방법론의 핵심이 되는 '재사용성','생산성' 등을 향상시키며 요구사항이 수시로 바뀌는 현대의 서비스 개발에 특화되어 빠르게 서비스를 피벗하거나 사용자가 급등할 수 있는 환경을 고려해 유연한 환경을 조성하는데 도움이 됩니다.

#### Uncontrolled Components

Uncontrolled Component는 상태를 직접 제어하지 않습니다.

```javascript 

const Hello = () => <div>Hello</div>

const Hi = () =>  <div>Hi</div>

```

#### Controlled component

상태를 직접제어하는 컴포넌트 입니다.

```javascript 

const HelloWorld = () => {
  const [value, setValue] = userState(0)

  return (
    <HelloComponent handleClick={setValue} >
      { value ? <Hello /> : <Hi /> }
    </HelloComponent>
  )
}                               

```

#### Pure component

Pure Component는 리액트를 이용한 컴포넌트 기본 개발 방법에서 가장 기본적인 컴포넌트입니다.   
Pure Component와 Component는 매우 유사하지만. 다른점이 있다면 React의 생명주기 메소드인 shouldComponentUpdate 를 다루는 방식에 있다.   
Component의 경우 setState가 실행되는 모든 순간에 Component가 Re-Rendering이 발생하지만, PureComponent의 경우에는 setState에서 Value가   
실제로 변경될 때만 Re-Rendering이 일어나게 된다.

```javascript

import React , { PureComponet } from 'react';

Class HelloWorld extend PureComponent {
  constructor(props){
    super(props);
    this.state = {
      value : 0
    }

    this.call_event = this.call_event.bind(this);
  }

  call_event() {=
    if(특정 조건){
      this.setState({
        counter: this.state .value + 변경값,
      });
    }else {
      this.setState({
        counter: this.state.value
      });
    }
  }

  render(){
    return (
      <div>
       {this.state.value}
        <br>
        <input
          type="button"
          onClick={this.call_event}
          value="click"
        />
      </div>
    )
  }
}


```

#### Portal component

React 버전 16부터 도입된 방법으로써, 컴포넌트를 렌더링할 때, UI를 어디에 렌더링 시킬지 DOM을 선택하여 부모 컴포넌트의 바깥에 렌더링할 수 있게
해주는 기능입니다. 기존의 리액트에서는 컴포넌트를 렌더링 하게 될 때, 하위 컴포넌트들은 부모 컴포넌트의 DOM 내부에서 동작했어야 했지만, Portals를 사용하면
DOM의 계층 구조 시스템에 종속되지 않으면서 컴포넌트를 렌더링 할 수 있습니다.

```javascript 

import ReactDOM from 'react-dom';

const ModalPortal = ({ children }) => {
  const el = document.getElementById('modal');
  return ReactDOM.createPortal(children, el);
};

export default ModalPortal;

```

```html 
                            
<body>
    You need to enable JavaScript to run this app.
    </noscript>
  <div id="root"><div>
  <div id="modal"><div>
<body>

```

우리가 해당 컴포넌트 방식을 사용하게 되면 우리가 원하는 컴포넌트의 결과를 DOM의 특정 엘리먼트에 렌더링 할 수 있게 됩니다.

#### Functional Component

함수형 컴포넌트라고 부르며, 16.8전가지는 함수형 컴포넌트와 클래스 기반 컴포넌트의 성능 차이는 거의 드러나지 않았습니다.
하지만 16.8에서 hooks가 도입되면서 함수형 컴포넌트에서 성능 최적화(useMemo, useCallback 등)를 진행할 수 있고, 라이프사이클(useEffect)를
제어할 수 있어 성능적으로 함수형 컴포넌트가 우위에 있다고 할 수 있습니다.

- 함수형 컴포넌트와 React Hook

    - useState
    - useEffect
    - useContext
    - useReducer
    - useCallback
    - useMemo
    - useRef
    - useImperativeHandle
    - useLayoutEffect
    - useDebugValue

#### Presentaitional Component

Presentaitional Component는 데이터와 관련된 이벤트 혹은 State 관리 없이 사용자에게 보여지는 뷰만을 관리합니다. DOM Element, style 등이 들어가며
ReadOnly Component라고도 부를 수 있습니다. 또한, 대부분의 경우 state를 갖고 있지 않으며, 갖고 있을 경우엔 데이터와 관련된 것이아니라 UI와 관련된것이어야 합니다.

#### Container Component

여러 컴포넌트를 관리하기 위해서 작성되며, DOM Elements를 직접적으로 수정하거나 관리하지 않습니다.

# ReactJS 활용

### if else 절

```javascript 


function List({ list }) {
  if (!list) {
    return null;
  }

  return (
    <div>
      {list.map(item => <ListItem item={item} />)}
    </div>
  );
}

function List({ list }) {
  if (!list) {
    return null;
  }

  if (!list.length) {
    return  <p>Sorry, the list is empty. </p>;
  } else {
    return (
       <div>
        {list.map(item =>  <ListItem item={item} />)}
       </div>
    );
  }
}
             

```

### 삼항 연산자

```javascript 

function Item({ item, mode }) {
  const isEditMode = mode === 'EDIT';

  return (
    <div>
      { isEditMode
        ?  <ItemEdit item={item} />
        :  <ItemView item={item} />
      }
    </div>
  );
}

function Item({ item, mode }) {
  const isEditMode = mode === 'EDIT';

  return (
    <div>
    {isEditMode ? (
        <ItemEdit item={item} />
    ) : (
        <ItemView item={item} />
    )}
    </div>
  );
}

```

### && 연산자

```javascript 

function LoadingIndicator({ isLoading }) {
  return (
     <div>
      { isLoading &&  <p>Loading... </p> }
     </div>
  );
}

```

### switch case 절

```javascript 

function Notification({ text, state }) {
  switch(state) {
    case 'info':
      return  <Info text={text} />;
    case 'warning':
      return  <Warning text={text} />;
    case 'error':
      return  <Error text={text} />;
    default:
      return null;
  }
}

Notification.propTypes = {
    text: React.PropTypes.string,
    state: React.PropTypes.oneOf(['info', 'warning', 'error'])
}  

function Notification({ text, state }) {
  return (
     <div>
      {(() => {
        switch(state) {
          case 'info':
            return  <Info text={text} />;
          case 'warning':
            return  <Warning text={text} />;
          case 'error':
            return  <Error text={text} />;
          default:
            return null;
        }
      })()}
     </div>
  );
}

```

### Enum을 이용한 처리

```javascript 

{% raw %}

const NOTIFICATION_STATES = {
  info:  <Info />,
  warning:  <Warning />,
  error:  <Error />,
};

function Notification({ state }) {
  return (
     <div>
      {NOTIFICATION_STATES[state]}
     </div>
  );
}

const getSpecificNotification = (text) => ({
  info:  <Info text={text} />,
  warning:  <Warning text={text} />,
  error:  <Error text={text} />,
});

function Notification({ state, text }) {
  return (
     <div>
      {getSpecificNotification(text)[state]}
     </div>
  );
}

function FooBarOrFooOrBar({ isFoo, isBar }) {
  const key = `${isFoo}-${isBar}`;
  return (
     <div>
      {{
        ['true-true']:  <FooBar />,
        ['true-false']:  <Foo />,
        ['false-true']:  <Bar />,
        ['false-false']: null,
      }[key]}
     </div>
  );
}

FooBarOrFooOrBar.propTypes = {
   isFoo: React.PropTypes.boolean.isRequired,
   isBar: React.PropTypes.boolean.isRequired,
}

{% endraw %}

```

### 다차원 연산 처리

```javascript 

function List({ list }) {
  const isNull = !list;
  const isEmpty = !isNull && !list.length;

  return (
     <div>
      { isNull
        ? null
        : ( isEmpty
          ?  <p>Sorry, the list is empty. </p>
          :  <div>{list.map(item =>  <ListItem item={item} />)} </div>
        )
      }
     </div>
  );
}

// Usage

 <List list={null} /> //  <div> </div>
 <List list={[]} /> //  <div> <p>Sorry, the list is empty. </p> </div>
 <List list={['a', 'b', 'c']} /> //  <div> <div>a </div> <div>b </div> <div>c </div> <div>


function List({ list }) {
  const isList = list && list.length;

  return (
       <div>
        { isList
          ?  <div>{list.map(item =>  <ListItem item={item} />)} </div>
          :  <NoList isNull={!list} isEmpty={list && !list.length} />
        }
       </div>
    );
  }

  function NoList({ isNull, isEmpty }) {
  return (!isNull && isEmpty) &&  <p>Sorry, the list is empty. </p>;
}

```

### ...을 이용하는 방식

```javascript 

const initialState = {
  data: {
    productQuantity: 0,
    installments: 0,
    totalPrice: 0,
    currencyId: 'USD',
    currencyFormat: '$'
  }
};

export default function(state = initialState, action) {
  switch (action.type) {
    case HELLOWORD :
      return {
        ...state,
        data: action.payload
      };
    default:
      return state;
  }
}

```

### Pagenation 구현하기

```javascript 

{% raw %}

var Paginator = React.createClass({
  _onClick(e) {
    setTimeout(function() { window.scrollTo(0, 0) }, 0)
  },

  render() {
    if (this.props.page === 1 && !this.props.hasNext) { return null }
    return <div className="Paginator">
      {this.props.page > 1 && <span className="Paginator__prev">
        <Link to={{pathname: `/${this.props.route}`, query: {page: this.props.page - 1}}} onClick={this._onClick}>Prev</Link>
      </span>}
      {this.props.page > 1 && this.props.hasNext && ' | '}
      {this.props.hasNext && <span className="Paginator__next">
        <Link to={{pathname: `/${this.props.route}`, query: {page: this.props.page + 1}}} onClick={this._onClick}>More</Link>
      </span>}
    </div>
  }
})

{% endraw %}

```

### env 설정 값을 사용하는 방식

```javscript

handleUpdate(update) {
  if (!this.isMounted()) {
    if (process.env.NODE_ENV !== 'production') {
      console.warn(
        `Skipping update as the ${this.props.type} Stories component is no longer mounted.`
      )
    }
    return
  }
  update.limit = update.ids.length
  this.setState(update)
}

```

### Javascript Filter의 활용 방식

```javascript 

const filter = ( props ) => {
  return props.valudList.filter( item => {
    if(item.isValid){
      return true;
    }else{
      return false;
    }
  }).slice(0, 10);
}

```

### 조건부 렌더링

```javascript 

const LoginControl = props => {

  const [ loginState, setLoginState] = useState({isLoggedIn : false})

  let handleLoginClick = function(){
    setLoginState({ isLoggedIn : true })
  }

  let handleLoginoutClick = function(){
    setLoginState({ isLoggedIn : false })
  }

  let appLayout = function(){

    if(isLoggedIn){
      button = <LogoutButton onClick={handleLogoutClick} />;
    }else{
      button = <LoginButton onClick={handleLoginClick} />;
    }

    return (
      <div>
        <Greeting isLoggedIn={isLoggedIn} />
        {button}
      </div>
    );
  }

  return appLayout;
}

ReactDOM.render(
  <LoginControl />,
  document.getElementById('root')
);

```

### 논리 && 연산자를 if를 인라인으로 표현하기

```javascript 

const MailBox = props => {
  const unreadMessages = props.unreadMessages;

  return (
    <div>
      <h1>Hello!</h1>
      {unreadMessages.length > 0 && 
        <h2>
          You have {unreadMessages.length} unread Messages. 
        </h2>
      }
    </dib>
  )
}

const messages = ['React', 'Re: React', 'Re:Re: React'];

ReactDOM.render(
  <MailBox unreadMessages={messages} />,
  document.getElementById('root')
)

```

### 리스트와 Key

```javascript 

const numbers = [ 1,2,3,4,5 ];
const doubled = numbers.map((number) => number * 2));
console.log(doubled);

```

- 여러개의 컴포넌트 렌더링 하기

```javascript 

const numbers = [ 1,2,3,4,5 ];
const listItems = numbers.map((number) => 
   <li>{number}</li>
);

ReactDOM.render(
  <ul>{listItems}</ul>,
  document.getElementById('root')
);

```

- 기본 리스트 컴포넌트

```javascript 

function NumberList(props){
  const numbers = props.numbers;
  // "key" 는 엘리먼트 리스트를 만들 때 포함해야 하는 특수한 문자열 어트리뷰트 입니다. 
  // Key는 React 가 어떤 항목을 변경, 추가 또는 삭제할 지 식별하는 것을 돕습니다. key 는 엘리먼트에 안정적인 고유성을 부여하기 위해 배열 내부의 엘리먼트를 지정해야 합니다. 

  /*
    Keys

    자식들이 key를 가지고 있다면, React는 key를 통해 기존 트리와 이후 트리의 자식들이 일치하는지 확인합니다. 
    예를 들어, key가 없는 리스트의 엘리먼트에 key를 추가하여 트리의 변환 작업이 효율적으로 수행되도록 
    수정할 수 있습니다. 

    Key는 배열 안에서 형제 사이에서 고유해야 하고 전체 범위에서 고유할 필요는 없습니다. 
    두 개의 다른 배열을 만들 때 동일한 key를 사용할 수 있습니다. 
  */
  const listItems = numbers.map((number) => 
    <li key={number.toString()} >{number}</li>
  );

  return (
    <ul>{listItems}</ul>
  );
}

const numbers = [1, 2, 3, 4, 5];
ReactDOM.render(
  <NumberList numbers={numbers} />,
  document.getElementById('root')
);

```

- JSX 에 map() 포함시키기

```javascript 

function NumberList(props){
  
  const numbers = props.numbers;

  const listItems = numbers.map((number) => 
    <ListItem key={number.toString()} value={number} />
  );

  return (
    <ul>
      {listItems}
    </ul>
  );
}

// JSX 를 사용하면 중괄호 안에 모든 표현식을 포함 시킬 수 있으므로 
// map() 함수의 결과를 인라인으로 처리할 수 있습니다. 
function NumberList(props) {
  const numbers = props.numbers;
  return (
    <ul>
      {numbers.map((numbers) => 
        <ListItem key={number.toString()} 
                     value={number} />  
      )}
    </ul>
  )
}

```

### Form

Javascript 함수로 폼의 제출을 처리하고 사용자가 폼에 입력한 데이터에 접근하도록 하는 것이 편리합니다. 이를 위한 편리한 표준 방식은 "제어 컴포넌트"라고 불리는 기술을 이용하는 것입니다.

```javascript 

class NameForm extens React.Component {
  constructor(props) {
    super(props);
    this.state = { value : ''};

    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);

  }

  handleChange(event) {
    this.setState( { value : event.target.value });
  }

  handleSubmit(event) {
    alert( ' A name was submitted:  ' + this.state.value );
    event.preventDefault();
  }

  /*
    value 어트리뷰트는 폼 엘리먼트에 설정되므로 표시되는 값은 항상 this.state.value 가 되고 React state는 신뢰가능한 단일 출처가 됩니다. 
    React state를 업데이트 하기 위해서 모든 키 입력에서 handleChange가 동작하기 때문에 사용자가 입력할 때 보여지는 값이 업데이트 됩니다. 
    제어 컴포넌트로 사용하면 모든 state 변화는 연관된 핸들러를 가집니다. 이것을 통해 사용자 입력을 수정하거나 유효성을 검사하는 것이 간단해집니다. 

    handleChange(event) {
      this.setState({value : event.target.value.toUpperCase()})
    }
  */
  render(){
    return (
      <form onSubmit={this.handleSubmit} >
        <lable>
          Name : 
            <input type="text" value={this.state.value} onChange={this.handleChange} />
        <label>
        <input type="submit"  value="Submit" />
      </form>
    );
  }
}                               


// React Hook을 활용한 입력값 바인딩 처리 방식 
import { useState } from "react";

const HelloForm = callback => {
  const [values, setValues] = useState({});

  const handleSubmit = event => {
    if (event) event.preventDefault();
    callback(values);
  };

  const handleChange = event => {
    event.persist();
    setValues(values => ({
      ...values,
      [event.target.name]: event.target.value
    }));
  };

  return {
    handleChange,
    handleSubmit,
    values
  };
};

export default HelloForm

```

### textarea 태그

```javascript 

let EssayForm = props => {
  const [ value, setValue ] = useState( {"value" : "Please write an essay about your favorite DOM element."});

  let handleChange = (event) => {
    setValue({ value : event.target.value });
  }

  let handleSubmit = (event) => {
    alert("An essay was submiited : " + value );
    event.preventDefault(); 
  }

  return (
    <form onSumbit={this.handleSumbit} >
      <label>
        Essay :
          <textarea value={value}
                        onChange={handleChange} />
          <input type="submit" value="Submit" />
      <label>
    <form>
  );
}

```

### select 태그

```javascript 

let FlavorForm = props => {
  let [ value, setValue ] = useState({value : 'coconut'});

  let handleChange = event => {
    setValue({ value : event.target.value });
  }

  let handleSubmit = event => {
    alert("Your favorite flavor is: " + value);
    event.preventDefault();
  }

  return {
    <form onSubmit={handleSubmit} >
      <label>
        Pick your favorite flavor :
        <select value={value} onChange={handleChange} >
          <option value="grapefruit" >Grapefruit</option>
          <option value="lime" >Lime</option>
          <option value="coconut" >Coconut</option>
          <option value="mango" >Mango</option>
        </select>
      </label>
      <input type="submit" value="Submit" />
    <form>

  }
}

/*
  select 태그에 multiple 옵션을 허용한다면 value 어트리 뷰트에 배열을 전달 할 수 있습니다. 

  <select multiple={true} value={['B', 'C']} >

*/

```

### file input

```javascript 


// 값이 읽기 전용이가 때문에 React에서는 비제어 컴포넌트입니다. 
<input type="file" />
// https://developer.mozilla.org/en-US/docs/Web/API/File/Using_files_from_web_applications
       

```

### 다중 입력 제어하기

```javascript 

// 여러 input 엘리먼트를 제어해야 할 때, 각 엘리먼트에 name 어트리뷰트를 추가하고 
// event.target.name 값을 통해 핸들러가 어떤 작업을 할지 선택할 수 있게 해줍니다. 

let Reservation = props => {

  const [ isGoing, setIsGoing ] = useState({isGoing : true});
  const [ numberOfGuests , setNumberOfGuests ] = useState({numberOfGuests : 2});

  let handleInputChange = event => {
    const target = event.target;
    const value = target.type === 'checkbox' ? target.checked : target.value;
    const name = target.name;

    switch(name){
      case "isGoing" : setIsGoing({isGoing: value}); break;
      case "numberOfGuests" : setNumberOfGuests({numberOfGuests: value}); break;
    }
  }

  return (
    <form>
      <label>
        <input 
          name="isGoing" 
          type="checkbox"
          checked={isGoing}
          onChange={handleInputChange}
      </label>
      <br /> 
      <label>
        <input
          name="numberOfGuests"
          type="number"
          value={numberOfGuests}
          onChange={handleChange} />
      </label>
    </form>
  )
}

```

### 제어되는 input Null 값

```javascript 

// 제어 컴포넌트에 value prop를 지정하면 의도하지 않는 한 사용자가 변경할 수 없습니다. 
// value로 설정했는데 여전히 수정할 수 있다면 실수로 value를 undefined 나 null 로 설정했을 수 있습니다. 

ReactDOM.render(<input value="hi" />, mountNode);

setTimeout(function() {
  ReactDOM.render(<input value={null} />, mountNode);
}, 1000)

```

### 컴포넌트가 렌더링 하는 것을 막기

```javascript 

import { useState } from react


let WarningBanner = prop => {
  if(!props.warn){
    return null;
  }

  return (
    <div className="warning" >
      Warning !
    <div>
  )
}

let Page = props => {
  const [ showWarning, setshowWarning ] = useState({showWarning : true})

  let handleToggleClick = function(){
    setshowWarning({
      showWarning : !showWarning
    })
  }

  let appLayout = (
    <div>
      <WarningBanner warn={handleToggleClick} >
      <button onClick={handleToggleClick} >
        { showWarning ? 'Hide' : 'Show' }
      </button>
    <div>
  );

  return appLayout;
}

ReactDOM.render(
  <Page />,
  document.getElementById('root')
)

```

### 조건부 연산자로 If-Else 구문을 인라인으로 표현하기

```javascript 

var appLayout = (
  <div>
    The user is <b> {isLoggedIn ? 'currently' : 'not' } <b> logged in.                         
  </div>
)

```

### Image Append 하기

Append 처리를 위해서 state를 사용하고, map을 이용하여 switch case 문을 적용함.

```javascript 

componentDidMount = index => {
  let wingsEndPointUrl = "http://localhost:8080/visit/factory/unit?requestUnit=Goliath";

  let indexVal = index;

  let unitListArray = this.state.unitList;

  console.log(indexVal);

  unitListArray.push(indexVal);

  this.setState({ unitList : unitListArray})

  axios.get(wingsEndPointUrl)
        .then(res => {
    const data = res.data;
    this.setState({ resourceGas : data.resourceGas, resourceMineral : data.resourceMineral });
  });

}

render() {
  const resourceGas  = this.state.resourceGas;
  const resourceMineral = this.state.resourceMineral;
 
  let Images = this.state.unitList.map(image  => {
    switch(image) {
      case 1:
        return <img className="w-50" src={siege_tank}></img>;
      case 2:
        return <img className="w-50" src={virture}></img>;
      case 3:
        return <img className="w-50" src={goliath}></img>;
      default:
        return null;
    }
   });


   <div className="col-10 mh-300"> 
      <div className="row" >
         <div className="col-5"><img src={factory}></img></div>
         <div className="col-2 overflow-auto">
              {Images}
           <img className="w-50" src={goliath}></img>
         </div>
         <div className="col-3"></div>
      </div>
  </div>

```

### Map을 이용한 다중 Element 및 이벤트 반환 방법

```javascript 

const someArr = ["A", "B", "C", "D"];

class App extends React.Component {
  state = {
    valueA: "",
    valueB: "some initial value",
    valueC: "",
    valueD: "blah blah"
  };

  handleChange = param => e => {
    const nextValue = e.target.value;
    this.setState({ ["value" + param]: nextValue });
  };

  render() {
    return (
      <div>
        {someArr.map(obj => {
          return (
            <div>
              <label>
                {`input ${obj}   `}
              </label>
              <input
                type="text"
                value={this.state["value" + obj]}
                onChange={this.handleChange(obj)}
              />
              <br />
              <br />
            </div>
          );
        })}
      </div>
    );
  }
}             

```

### For 문을 이용한 중첩 Component 호출 처리

```javascript 

class AppComponent extends React.Component {
  state = {
    numChildren: 0
  }

  render () {
    const children = [];

    for (var i = 0; i < this.state.numChildren; i += 1) {
      children.push(<ChildComponent key={i} number={i} />);
    };

    return (
      <ParentComponent addChild={this.onAddChild}>
        {children}
      </ParentComponent>
    );
  }

  onAddChild = () => {
    this.setState({
      numChildren: this.state.numChildren + 1
    });
  }
}

const ParentComponent = props => (
  <div className="card calculator">
    <p><a href="#" onClick={props.addChild}>Add Another Child Component</a></p>
    <div id="children-pane">
      {props.children}
    </div>
  </div>
);

const ChildComponent = props => <div>{"I am child " + props.number}</div>;

```

### state 끌어올리기

React 애플리케이션 안에서 변경이 일어나는 데이터에 대해서는 "진리의 원천"을 하나만 두어야 합니다. 보통의 경우, state는 렌더링에 그 값을 필요로 하는 컴포넌트에 먼저 추가됩니다. 그리고 나서 다른 컴포넌트도 역시 그 값이 필요하게 되면 그 값을 그들의 가장 가까운 공통 조상으로 끌어올리면 됩니다. 다른 컴포넌트 간에 존재하는 state를 동기화 시키려고 노력하는 대신 하향식 데이터 흐름에 기대는 걸 추천합니다

```javascript 

class Calculator extends React.Component {
  constructor(props) {
    super(props);
    this.handleCelsiusChange = this.handleCelsiusChange.bind(this);
    this.handleFahrenheitChange = this.handleFahrenheitChange.bind(this);
    this.state = {temperature: '', scale: 'c'};
  }

  handleCelsiusChange(temperature) {
    this.setState({scale: 'c', temperature});
  }

  handleFahrenheitChange(temperature) {
    this.setState({scale: 'f', temperature});
  }

  render() {
    const scale = this.state.scale;
    const temperature = this.state.temperature;
    const celsius = scale === 'f' ? tryConvert(temperature, toCelsius) : temperature;
    const fahrenheit = scale === 'c' ? tryConvert(temperature, toFahrenheit) : temperature;

    return (
      <div>
        <TemperatureInput
          scale="c"
          temperature={celsius}
          onTemperatureChange={this.handleCelsiusChange} />

        <TemperatureInput
          scale="f"
          temperature={fahrenheit}
          onTemperatureChange={this.handleFahrenheitChange} />

        <BoilingVerdict
          celsius={parseFloat(celsius)} />

      <div>
    );
  }
}

class TemperatureInput extends React.Component {
  constructor(props) {
    super(props);
    this.handleChange = this.handleChange.bind(this);
  }

  handleChange(e) {
    this.props.onTemperatureChange(e.target.value);
  }

  render() {
    const temperature = this.props.temperature;
    const scale = this.props.scale;
    return (
      <fieldset>
        <legend>Enter temperature in {scaleNames[scale]}:<legend>
        <input value={temperature}
               onChange={this.handleChange} />
      <fieldset>
    );
  }
}

```
