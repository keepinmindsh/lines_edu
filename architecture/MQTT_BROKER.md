# MQTT Broker  

MQTT는 1999년에 발표한 오픈 프로토콜로 낮은 대역폭, 높은 지연이나 신뢰할 수 없는 네트워크를 위하여 설계된 경량적인 메시지 프로토콜이다. 

## MQTT - Hello World 

#### Mac 에 직접 설치하는 방법으로 테스트 진행 

- [https://mosquitto.org/download/](https://mosquitto.org/download/)

```shell
$ brew install mosquitto
```


#### 아래의 문법은 동작하지 않음

- conf 파일도 넣어줬으나 제대로 동작하지 않는 이슈가 있음, 추가 다른 설정이 필요해보이는데 확인하지 못함. 

```shell
# make volume 
$ mkdir -p /Users/lines/mos-docker/mosquitto && mkdir -p /Users/lines/mos-docker/mosquitto/log && mkdir -p /Users/lines/mos-docker/mosquitto/data

$ docker run -it --name lines-mos-mqtt -p 1883:1883  -v /Users/lines/mos-docker/mosquitto:/mosquitto/ -v /Users/lines/mos-docker/mosquitto/log -v /Users/lines/mos-docker/mosquitto/data:/mosquitto/data -v /Users/lines/mos-docker/mosquitto/config:/mosquitto/config  eclipse-mosquitto
```