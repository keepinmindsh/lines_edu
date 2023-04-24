# Test-Driven Development

Test Driven Development(TDD)는 소프트웨어를 빌드하기 위한 기술로, 소프트웨어 개발 방식을 테스트를 기반으로 하여 진행합니다.  
이는 Extreme Programming의 일환으로 켄트 벡에 의해 1990년대 후반에 개발되었습니다.  

테스트 주도 개발을 하는 가장 기본적인 절차는 아래와 같습니다.  

1. 추가하려는 기능에 대해서 테스트 코드를 작성합니다. 
2. 테스트가 통과되기 전까지 코드를 작성합니다. 
3. 신규 코드와 오래된 코드를 잘 구조화될 수 있도록 작성합니다. 

위의 3가지 단계에 대한 순환 작업을 통해서 요구사항에 대한 기능을 누적시킵니다. 여기에서 가장 중요한 주안점은 요구사항에 대한 실제 구현 코드로 시작하는 것이 아닌 
테스트 코드로 부터 실제 구현 코드가 작성되어 간다는 것입니다.  

이는 요구사항에 대해서 명확하게 구현할 함수 외형을 테스트로 작성하면서 실제 구현해야할 기능에만 초점을 맞출 수 있습니다.     
테스트를 우선하여 생각하게 되면, 구현되지 않는 코드에서 시작해야하기 때문에 Mocking 등의 바탕이 되면 인터페이스에 대해서 당연히 고민할 수 밖에 없습니다.  
이는 자연스러운 추상화 및 코드를 구조적으로 작성하는 능력을 키워주게 됩니다.  

**마지막으로 가장 중요한 것은** 절대 1번, 2번 을 진행후 3번을 진행하는 것에 대해서 소홀히 해서는 안된다는 것입니다. 1번, 2번, 3번의 과정이 짧은 주기로 반복적으로 작성되어야 
정말 올바른 코드가 작성되게 되고, 그렇게 작성된 코드는 가독성, 확장성을 위한 추상화 등이 충분히 고려된 좋은 코드가 됩니다.  



> [TDD](https://martinfowler.com/bliki/TestDrivenDevelopment.html) 