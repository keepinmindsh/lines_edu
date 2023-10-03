

## Class Property 

Class 에서 생성한 오브젝트를 사용하는 방법은 아래의 코드를 참고할 것! 

```python
class ConcreteBuilder1(Builder):

    def __init__(self) -> None:
        self.reset()

    def reset(self) -> None:
        self._product = Product1()

    @property
    def dropship(self) -> Product1:
        product = self._product
        self.reset()
        return product

    def get_on_the_ship(self) -> None:
        self._product.add("PartA1")

    def get_off_the_ship(self) -> None:
        self._product.add("PartB1")

if __name__ == "__main__":

    director = Director()
    builder = ConcreteBuilder1()
    director.builder = builder

    print("Standard full featured product: ")
    builder.dropship.list_parts()
```

## Keyword 

### pass 

- pass는 아무것도 수행하지 않는 문법으로, 임시로 코드를 작성할 때 주로 사용한다.

```python
class Builder(ABC):

    @property
    @abstractmethod
    def dropship(self) -> None:
        pass
```
