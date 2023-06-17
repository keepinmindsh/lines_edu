# MQTT Broker  

MQTT는 1999년에 발표한 오픈 프로토콜로 낮은 대역폭, 높은 지연이나 신뢰할 수 없는 네트워크를 위하여 설계된 경량적인 메시지 프로토콜이다. 

## MQTT - Hello World 

### Mac 에 직접 설치하는 방법으로 테스트 진행 

- [https://mosquitto.org/download/](https://mosquitto.org/download/)

```shell
$ brew install mosquitto
```

```shell
brew services start mosquitto

brew services stop mosquitto 

brew services restart mosquitto
```

> [Mosquitto 설치 및 실행](https://velog.io/@aimzero9303/Mac-MQTT-mosquitto-%EC%84%A4%EC%B9%98-%ED%85%8C%EC%8A%A4%ED%8A%B8)

### 실제 호출 테스트

```shell
# Consumer Listen - 다중으로 호출 및 세팅이 가능하다. 
$ /opt/homebrew/Cellar/mosquitto/2.0.15/bin/mosquitto_sub -h localhost -p 1883 -t every
```

```shell
# Producer Call - Message를 다중으로 호출한다. 
$ /opt/homebrew/Cellar/mosquitto/2.0.15/bin/mosquitto_pub -h localhost -p 1883 -t every -m "테스트 메세지 호출한다이~다중 호출 가능하네?"
```

위의 코드에 대한 정상 호출 확인 완료하였음. Brew 내에 MQTT Broker 를 설정하는 방식으로 처리

### 아래의 문법은 동작하지 않음

- conf 파일도 넣어줬으나 제대로 동작하지 않는 이슈가 있음, 추가 다른 설정이 필요해보이는데 확인하지 못함. 

```shell
# make volume 
$ mkdir -p /Users/lines/mos-docker/mosquitto && mkdir -p /Users/lines/mos-docker/mosquitto/log && mkdir -p /Users/lines/mos-docker/mosquitto/data

$ docker run -it --name lines-mos-mqtt -p 1883:1883  -v /Users/lines/mos-docker/mosquitto:/mosquitto/ -v /Users/lines/mos-docker/mosquitto/log -v /Users/lines/mos-docker/mosquitto/data:/mosquitto/data -v /Users/lines/mos-docker/mosquitto/config:/mosquitto/config  eclipse-mosquitto
```


## MQTT Go Client Library

- [Client Connection Test Sample](https://github.com/eclipse/paho.golang/blob/master/paho/client_test.go)