# Template Method 

## Basic Information

- 패턴형식 : 행위패턴 

## Gof's Description

객체의 연산이 알고리즘의 뼈대 만들 정의하고 각 단계에서 수행할 구체적 처리는 서브 클래스 쪽에서 미룹니다.   
알고리즘의 구조 자체는 그대로 놔둔 채 알고리즘 각 단계의 처리를 서브 클래스에서 처리할 수 있게 합니다.

## 동기

예를 들어, 여러분이 어떤 프로그램을 개발한다고 해보겠습니다. 해당 프로그램이 하는 역할은 ToDoList 라고 하면, ToDoList는 

내가 해야할 일을 등록하는 과정이 있습니다.   

- 제목 입력 
- 내용 입력
- 기간 입력 
- 저장 처리 

위의 각 4가지 과정을 처리한다고 하겠습니다.   

그 뒤에 서비스가 확장되었습니다. ToDoList가 Health에 대해서만 추가로 처리해야하는 프로세스가 있다고 할때, 

- 제목 입력
- 헬스를 위한 목표 입력칸 추가 
- 기간 입력 
- 저장 처리 

위와 같이 4가지 단계를 처리한다고 할 때, 첫번째 항목과 다른 점은 내용을 입력하는 방식이 된다고 하면, 

이때 흔히 우리는 첫번째 ToDoList의 항목을 그대로 둔채, 추가로 헬스에 대한 내용을 작성하게 되는데,    
여기에서 동일한 로직에 대한 중복 코드가 발생하기 시작합니다.   

이 상황에서 적용하기 좋은 것이 Template Method 패턴입니다.   

## 활용 

- 어떤 한 알고리즘을 이루는 부분 중 변하지 않는 부분을 한번 정의해 놓고 다양해질 수 있는 부분을 서브클래스에서 정의할 수 있도록 남겨주고자 할 때, 
- 서브 클래스 사이의 공통 적인 행동을 추철하여 하나의 공통 클래스에 몰아둠으로써 코드의 중복을 피하고 싶을 때, 
- 서브 클래스의 확장을 고려할 때, 

## 코드 예시 

### Golang 

```go 
package main

import (
	"todolist/app/todo"
	"todolist/domain"
)

func main() {
	healthTodo := todo.MustNewTodo(domain.HealthToDo)

	healthTodo.MakeTitle()
	healthTodo.MakeContent()
	healthTodo.MakePeriod()
	healthTodo.MakeConfirm()

	taskTodo := todo.MustNewTodo(domain.TaskTodo)

	taskTodo.MakeTitle()
	taskTodo.MakeContent()
	taskTodo.MakePeriod()
	taskTodo.MakeConfirm()
}

```

공통적인 역할과 개별로 처리해야할 역할을 분리

```go 
package domain

type ToDoType string

const (
	HealthToDo ToDoType = "Health"
	TaskTodo   ToDoType = "Task"
)

type ToDo interface {
	CustomToDo
	CommonToDo
}

type CustomToDo interface {
	MakeContent()
}

type CommonToDo interface {
	MakeTitle()
	MakePeriod()
	MakeConfirm()
}
```

Template 메소드를 구조체를 이용하여 정의

```go 
package todo

import (
	"todolist/app/usecase"
	"todolist/domain"
)

type ToDo struct {
	domain.CommonToDo
	domain.CustomToDo
}

func MustNewTodo(doType domain.ToDoType) ToDo {
	switch doType {
	case domain.TaskTodo:
		return ToDo{
			usecase.NewCommonToDo(),
			usecase.NewTaskToDo(),
		}
	case domain.HealthToDo:
		return ToDo{
			usecase.NewCommonToDo(),
			usecase.NewHealthTodo(),
		}
	default:
		panic("error")
	}
}
```

```go 
package usecase

import (
	"fmt"
	"todolist/domain"
)

type Common struct {
}

func (c Common) MakeTitle() {
	fmt.Println("타이틀을 작성합니다.")
}

func (c Common) MakePeriod() {
	fmt.Println("기간을 입력합니다.")
}

func (c Common) MakeConfirm() {
	fmt.Println("확인 및 저장힙니다.")
}

func NewCommonToDo() domain.CommonToDo {
	return &Common{}
}
```

```go 
package usecase

import (
	"fmt"
	"todolist/domain"
)

type HealthTodo struct {
}

func (h HealthTodo) MakeContent() {
	fmt.Println("운동을 위한 상세 운동 방법을 입력합니다.")
}

func NewHealthTodo() domain.CustomToDo {
	return &HealthTodo{}
}
```

```go 
package usecase

import (
	"fmt"
	"todolist/domain"
)

type TaskTodo struct {
}

func (t TaskTodo) MakeContent() {
	fmt.Println("오늘의 할일을 간단하게 작성합니다.")
}

func NewTaskToDo() domain.CustomToDo {
	return &TaskTodo{}
}
```

### Java 

```java 
public class Caller {

  public static void main(String[] args) {
      DBTemplate dbTemplate = new OracleDB();

      List<Map> resutList = (List<Map>)dbTemplate.execute("SELECT * FROM TB_ZZ_USER");

      resutList.forEach(value -> {
          System.out.println(value.get("USER_ID"));
      });
  }
}
```

```java 
public abstract class DBTemplate {

  public Object selectQuery(String sql){
      return execute(sql);
  }

  public List<Map> execute(String queryStatment){

      ResultSet resultSet = null;
      List<Map> resultList = null;
      Connection connection = null;
      PreparedStatement preparedStatement = null;

      try {

          String[] connectionInfo = getDBInformation();

          connection = getConnection(connectionInfo);

          preparedStatement = executeStatement(connection, queryStatment);

          resultSet = getResultSet(preparedStatement);

          resultList = getResultMapRows(resultSet);

      }catch (Exception ex){
          ex.printStackTrace();
      }finally {
          try{
              if(connection != null ) releaseConnection(connection);
              if(preparedStatement != null ) releasePrepareStatement(preparedStatement);
              if(resultSet != null ) releaseResultSet(resultSet);
          }catch (Exception ex){
              ex.printStackTrace();
          }

      }

      return resultList;
  }

  /**
    * ResultSet을 Row마다 Map에 저장후 List에 다시 저장.
    * @param rs DB에서 가져온 ResultSet
    * @return Listt<map> 형태로 리턴
    * @throws Exception Collection
    */
  private List<Map> getResultMapRows(ResultSet rs) throws Exception
  {
      // ResultSet 의 MetaData를 가져온다.
      ResultSetMetaData metaData = rs.getMetaData();
      // ResultSet 의 Column의 갯수를 가져온다.
      int sizeOfColumn = metaData.getColumnCount();

      List<Map> list = new ArrayList<Map>();
      Map<String, Object> map;
      String column;
      // rs의 내용을 돌려준다.
      while (rs.next())
      {
          // 내부에서 map을 초기화
          map = new HashMap<String, Object>();
          // Column의 갯수만큼 회전
          for (int indexOfcolumn = 0; indexOfcolumn < sizeOfColumn; indexOfcolumn++)
          {
              column = metaData.getColumnName(indexOfcolumn + 1);
              // map에 값을 입력 map.put(columnName, columnName으로 getString)
              map.put(column, rs.getString(column));
          }
          // list에 저장
          list.add(map);
      }
      return list;
  }

  private void releaseConnection(Connection connection) throws SQLException {
      connection.close();
  }

  private void releaseResultSet(ResultSet resultSet) throws SQLException {
      resultSet.close();
  }

  private void releasePrepareStatement(PreparedStatement preparedStatement) throws SQLException {
      preparedStatement.close();
  }

  public abstract String[] getDBInformation();

  public abstract Connection getConnection(String[] connectionInfo) throws SQLException;

  public abstract PreparedStatement executeStatement(Connection connection, String sql) throws SQLException;

  public abstract ResultSet getResultSet(PreparedStatement preparedStatement) throws SQLException;

}

public class OracleDB extends DBTemplate {
  @Override
  public String[] getDBInformation() {
      String jdbcDriver = "jdbc:oracle:thin:@***.***.***.***:***/*****";
      String dbUser = "******";
      String dbPassword = "******";

      String[] resultInfo = {jdbcDriver, dbUser, dbPassword};

      return resultInfo;
  }

  @Override
  public Connection getConnection(String[] connectionInfo) throws SQLException {
      return DriverManager.getConnection(connectionInfo[0], connectionInfo[1], connectionInfo[2]);
  }

  @Override
  public PreparedStatement executeStatement(Connection connection, String sql) throws SQLException {
      return connection.prepareStatement(sql);
  }

  @Override
  public ResultSet getResultSet(PreparedStatement preparedStatement) throws SQLException {
      return preparedStatement.executeQuery();
  }
}
```