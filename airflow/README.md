# Index 

- [AirFlow](#airflow)
  - [AirFlow 소개](#airflow-is)
  - [AirFlow 의 좋은 점](#airflows-good-things)
- [AirFlow 쿠버네티스 환경에 구성하기](#install-airflow-on-kubernetes)
  - [쿠버네티스 환경의 AirFlow 구성 및 동작 검증](#valuesyaml-of-kubernetes-applied)
- [Docker Desktop에서 Airflow 쿠버네티스 환경 구성하기](#install-airflow-on-docker-desktop)
- [AirFlow Study](#airflow-study)
  - [Airflow GetStarted](#get-started)
  - [AirFlow Docs에서 참고할 만한 사항](#airflow-documentation-help)
    - [Airflow Docker Image](#airflow-docker-image)
    - [Airflow 를 쿠버네티스에서 Helm 차트를 이용해 설치하기](#airflow-on-kubernetes-with-heml-chart)
    - [Airflow 사전 설치 환경](#prerequisites-of-airflow)
    - [Airflow 의 라이프 사이클 및 EOL](#lifecycle-and-system-eol-for-airflow)
    - [Airflow 에 대한 갱신 및 마이그레이션](#upgrade-and-migrate-for-airflow)
    - [Airflow 을 위한 보안 구성](#security-for-airflow)
    - [Airflow RateLimiting 설정](#setting-of-airflow-ratelimiting)
    - [민감한 정보에 대한 Masking 처리](#handling-masking-on-sensitive-information)
  - [Airflow 기본](#airflow-basic)

---

# AirFlow 

- [https://airflow.apache.org/](https://airflow.apache.org/)

## Airflow Is

Airbnb에서 개발했으며 데이터 파이프라인을 관리하고 오케스트레이션 하기 위한 강력한 오픈 소스 플랫폼으로, 동작 파이프라인 생성, 작업 재시도 및 모니터링과 같은 주요 기능을 제공함.

![Airflow Images](https://github.com/lines-code/lines-workflow-airflow-python/blob/master/docs/assets/AirflowConcepts.png)

### Scheduler 
- 예약된 워크플로우를 트리거하고, Executor에게 작업을 전달하고 실행

### Executor 
- 실행 중인 작업을 처리하는 실행 프로그램 
- 기본적으로 Airflow 를 설치하면 스케줄러 내부의 모든 것을 실행하지만, 
- Production 을 위한 Airfow 작업에서는 Worker에게 작업을 전달합니다.

Executor 란 작업자들에게 작업을 실행 시키는 역할을 한다. 종류로는 SequentialExecutor, CeleryExecutor, KubernetesExecutor 등이 있다.

#### Executors 

- Sequential Executors
  - default executor 이다. 
  - 클러스터를 구성하지 않고 작업은 순차적으로 진행된다. 
  - scheduler와 같이 진행된다. 
  - DB는 보통 Sqlite를 사용한다. 
  - 쉽고 간결하다 
  - 스케일 인아웃 불가능 
  - 순차실행이라 병렬처리 불가능 
  - 테스트 디버깅용으로 좋다. 
- Local Executor 
  - scheduler와 함께 동작한다. 
  - 병렬 처리 가능 
  - 쉽게 세팅 가능 
  - 가볍고 싸다 
  - 병렬처리 가능하다. 
  - 스케일 인아웃을 할 수 없다. 
  - 테스트 디버깅 용으로 좋다. 작은 스케일에서 사용하기 좋다. 
- Celery Executor 
  - rabbitMQ를 사용하여야 한다. 
  - 분산 태스크 처리가 가능하다. 
  - 수평 구조로 스케일링 할 수 있다. 
  - 내 결함성을 제공한다. 
  - 자원 낭비가 있을 수 있다. 
  - 비용 효율적이지 않다. 
  - production에 적합하다.
  - worker가 항상 동작하고 있어야 한다. 
- Kubernetes Executor 
  - Kubernetes API를 사용하여 매니저 된다. 
  - 스케일러블 하기 쉽다. 
  - 내 결함성을 가진다. 
  - 자원을 개별적으로 태스크 돌릴 수 있다. 
  - production에 적합하다. 
  - 워커를 필요할 때 사용이 가능하다. 

#### 사용 가이드 라인 

Celery Executor는 상용버전에는 적합하나 Rabbimq나 Redis의 큐나 세션을 이용하여 Executor의 이중 작업을 꼭 방지 해야한다. 작업이 두번 실행 되면 안되기때문이다. 그리고 Celery에 경우 클러스터를 구축하여야 하기 때문에 Zookeeper나, etcd, consul을 사용하여 자기 자신이 replication을 하는 클러스터를 구축을 해야하는데, web ui, scheduler, rabbimq, 클러스터 broker3대, DB서버 등을 계산한다면 컴퓨터가 최소 4~5대 정도는 필요하다. orchestration툴에 5대의 컴퓨터는 조금 오버 스펙이라고 생각하여 잘 사용하지는 않는 것 같다. 제대로 관리 하고 싶다면 사용해보자.

Kubernetes의 경우에는 Container를 만들고 작업을 실행한다음 바로 삭제 하기 때문에 Log를 제한적으로 볼수 밖에 없다. 그래서 K8s에 대해 다 아는 분이 꼭 사용하여야 한다. Celery보다 자유로운 구성을 할 수 있어 Celery보다 resource를 효율적으로 사용할 수 있다. 하지만 러닝커브가 큰 k8s를 잘 다뤄야 한다는 문제점이 있고, Task를 수행한 Executor를 바로 삭제하기 때문에 Log가 많은 정보를 담고 있지 않다.

##### 참조 

> [https://magpienote.tistory.com/225](https://magpienote.tistory.com/225)

### DAG Directory 
- /dags 정의한 DAG 파일을 포함하는 폴더
- Scheduler와 Executor, 그리고 Executor가 가진 모든 Worker 가 읽는 폴더입니다.
- Dag 내의 Task 실행 단위에서 핸들링 할 수 있는 Operator 입니다. 
  - https://airflow.apache.org/docs/apache-airflow/stable/core-concepts/operators.html

### Metadata Database 
- scheduler, executor 및 webserver가 상태를 저장하는 데 사용하는 데이터베이스

### Webserver 
- 사용자의 편의성을 위해 DAG와 Task의 동작을 검사하고 트리거하고, 디버깅할 수 있는 UI를 제공하는 웹 서버

```python 
import textwrap
from datetime import datetime, timedelta

# The DAG object; we'll need this to instantiate a DAG
from airflow.models.dag import DAG

# Operators; we need this to operate!
from airflow.operators.bash import BashOperator
with DAG(
    "tutorial_test_jsh",
    # These args will get passed on to each operator
    # You can override them on a per-task basis during operator initialization
    default_args={
        "depends_on_past": False,
        "email": ["airflow@example.com"],
        "email_on_failure": False,
        "email_on_retry": False,
        "retries": 1,
        "retry_delay": timedelta(minutes=5),
        # 'queue': 'bash_queue',
        # 'pool': 'backfill',
        # 'priority_weight': 10,
        # 'end_date': datetime(2016, 1, 1),
        # 'wait_for_downstream': False,
        # 'sla': timedelta(hours=2),
        # 'execution_timeout': timedelta(seconds=300),
        # 'on_failure_callback': some_function, # or list of functions
        # 'on_success_callback': some_other_function, # or list of functions
        # 'on_retry_callback': another_function, # or list of functions
        # 'sla_miss_callback': yet_another_function, # or list of functions
        # 'trigger_rule': 'all_success'
    },
    description="A simple tutorial DAG",
    schedule=timedelta(days=1),
    start_date=datetime(2021, 1, 1),
    catchup=False,
    tags=["hello_world"],
) as dag:

    # t1, t2 and t3 are examples of tasks created by instantiating operators
    t1 = BashOperator(
        task_id="helloworld",
        bash_command="date",
    )

    t2 = BashOperator(
        task_id="hi1",
        depends_on_past=False,
        bash_command="sleep 5",
        retries=3,
    )
    t1.doc_md = textwrap.dedent(
        """\
    #### Task Documentation
    You can document your task using the attributes `doc_md` (markdown),
    `doc` (plain text), `doc_rst`, `doc_json`, `doc_yaml` which gets
    rendered in the UI's Task Instance Details page.
    ![img](http://montcs.bloomu.edu/~bobmon/Semesters/2012-01/491/import%20soul.png)
    **Image Credit:** Randall Munroe, [XKCD](https://xkcd.com/license.html)
    """
    )

    dag.doc_md = __doc__  # providing that you have a docstring at the beginning of the DAG; OR
    dag.doc_md = """
    This is a documentation placed anywhere
    """  # otherwise, type it like this
    templated_command = textwrap.dedent(
        """
    {% for i in range(5) %}
        echo "{{ ds }}"
        echo "{{ macros.ds_add(ds, 7)}}"
    {% endfor %}
    """
    )

    t3 = BashOperator(
        task_id="hi2",
        depends_on_past=False,
        bash_command=templated_command,
    )

    t4 = BashOperator(
        task_id="hi3",
        depends_on_past=False,
        bash_command=templated_command,
    )

    t5 = BashOperator(
        task_id="hi4",
        depends_on_past=False,
        bash_command=templated_command,
    )

    t6 = BashOperator(
        task_id="h5",
        depends_on_past=False,
        bash_command=templated_command,
    )

    t7 = BashOperator(
        task_id="h6",
        depends_on_past=False,
        bash_command=templated_command,
    )

    t1 >> [t2, t3, t7] >> t5 >> t6 >> t4
```


## Airflow's Good Things

- Kubernetes Executor를 통해서 Pod 단위의 배포 및 처리가 가능함. 
  - Kubernetes Executor를 사용하면 Worker를 Pod 형태로 동적으로 생성하게 됩니다. 
  - DAG이 실행될 때 Task 하나당 하나의 Worker Pod가 배포 및 실행된 후 삭제됩니다. 
  - 실행할 DAG이 없는 경우 Kubernetes 위에 Worker Pod이 존재하지 않습니다.

- GigSync는 주기적으로 Airflow DAG 폴더가 주기적으로 업데이트 됩니다.
  - Workflow 코드에 대한 이력 관리가 가능하다. 
  - Dag Generator를 API를 통해서 전달받을 수 있는 방식으로 구성한 뒤 해당 소스를 git repository로 Sync 처리하면 코드 작성후 사용할 수 있는 방식으로 처리 가능함. 
  - 각 필요에 따라 자체적인 템플릿 구성 및 생성이 가능함.

```yaml 
gitSync:
    enabled: true
    repo:  ${Organization-ID}@github.com:${Orgarnizations}/${Project-Repo-Name}.git  # 조직의 Repo 연결 
    branch: master
    rev: HEAD
    ref: master
    depth: 1
    maxFailures: 0
    subPath: ""
    sshKeySecret: airflow-git-ssh-secret
    period: 5s
    wait: ~
    containerName: git-sync
    uid: 65533
```

- Web에서 원격으로 제어 가능한 API 제공 
  - 각 Dag를 개발자/운영자 입맛에 맞게 제어 및 관리가 가능해요.  
  - https://airflow.apache.org/docs/apache-airflow/stable/stable-rest-api-ref.html#tag/DAG

- Airbnb에서 개발했으며 데이터 파이프라인을 관리하고 오케스트레이션 하기위한 강력한 오픈 소스 플랫폼으로, 동작 파이프라인 생성, 작업 재시도 및 모니터링과 같은 주요 기능을 제공함.

# Install Airflow on Kubernetes 

```shell 
gcloud config set project ${project-id}

gcloud container clusters get-credentials ${cluster-id} --region us-central1-f --project ${project-id}

kubectl create namespace airflow && kubectl config set-context --current --namespace=airflow

helm repo add apache-airflow https://airflow.apache.org
helm repo update
helm search repo airflow
helm install airflow apache-airflow/airflow --namespace airflow --debug

Airflow Webserver:     kubectl port-forward svc/airflow-webserver 8080:8080 --namespace airflow
Default Webserver (Airflow UI) Login credentials:
    username: admin
    password: admin
Default Postgres connection credentials:
    username: postgres
    password: postgres
    port: 5432


helm show values apache-airflow/airflow > values.yaml

# values.yaml 파일에서 기본 executor를 KubernetesExecutor로 변경

helm ls -n airflow


# Register Public Key 
ssh-keygen -t rsa -b 4096 -C "****@****.**"

# 잘안됨...
# github 내에 Deploy Key 등록 처리 
# https://www.base64encode.org/ 통해서 인코딩 
#kubectl apply -f airflow-ssh-secret.yaml

# 다른 방법
# https://kimjinung.tistory.com/5
kubectl create secret generic airflow-git-ssh-secret \
  --from-file=gitSshKey=/Users/lines/sources/03_onlines_corp/settings/airflow_settings/id_rsa \
  --from-file=id_ed25519.pub=/Users/lines/sources/03_onlines_corp/settings/airflow_settings/id_rsa.pub \
  -n airflow

helm show values apache-airflow/airflow > values.yaml


helm upgrade --install airflow apache-airflow/airflow -n airflow -f values.yaml --debug

# https://kubernetes.io/ko/docs/tutorials/stateless-application/expose-external-ip-address/
kubectl get deployments -A
kubectl expose deployment airflow-webserver --type=LoadBalancer --name=airflow-service

# https://airflow.apache.org/docs/apache-airflow/2.1.1/_api/airflow/sensors/sql/index.html
```

## values.yaml of kubernetes applied

- airflow config 중의 일부 

```shell 
  gitSync:
    enabled: true

    # git repo clone url
    # ssh example: git@github.com:apache/airflow.git
    # https example: https://github.com/apache/airflow.git
    repo: ${Organization-ID}@github.com:${Orgarnizations}/${Project-Repo-Name}.git
    branch: master
    rev: HEAD
    # The git revision (branch, tag, or hash) to check out, v4 only
    ref: master
    depth: 1
    # the number of consecutive failures allowed before aborting
    maxFailures: 0
    # subpath within the repo where dags are located
    # should be "" if dags are at repo root
    subPath: ""
    # if your repo needs a user name password
    # you can load them to a k8s secret like the one below
    #   ---
    #   apiVersion: v1
    #   kind: Secret
    #   metadata:
    #     name: git-credentials
    #   data:
    #     # For git-sync v3
    #     GIT_SYNC_USERNAME: <base64_encoded_git_username>
    #     GIT_SYNC_PASSWORD: <base64_encoded_git_password>
    #     # For git-sync v4
    #     GITSYNC_USERNAME: <base64_encoded_git_username>
    #     GITSYNC_PASSWORD: <base64_encoded_git_password>
    # and specify the name of the secret below
    #
    # credentialsSecret: git-credentials
    #
    #
    # If you are using an ssh clone url, you can load
    # the ssh private key to a k8s secret like the one below
    #   ---
    #   apiVersion: v1
    #   kind: Secret
    #   metadata:
    #     name: airflow-ssh-secret
    #   data:
    #     # key needs to be gitSshKey
    #     gitSshKey: <base64_encoded_data>
    # and specify the name of the secret below
    sshKeySecret: airflow-git-ssh-secret
```

# Install Airflow on Docker Desktop 

- 기본적인 Docker Desktop에 내장되어 있는 Kubernetes 설치 

설치된 kubernetes의 nodes 점검

```shell 
kubectl get nodes
NAME             STATUS   ROLES           AGE   VERSION
docker-desktop   Ready    control-plane   22m   v1.29.1
```


```shell 
kubectl create namespace airflow && kubectl config set-context --current --namespace=airflow

# Airflow 점검 
kubectl get namespaces

helm repo add apache-airflow https://airflow.apache.org
helm repo update
helm search repo airflow
helm install airflow apache-airflow/airflow --namespace airflow --debug

# id_rsa, id_rsa.pub 생성
ssh-keygen -t rsa -b 4096 -C "keepinmindsh@gmail.com"

# id_rsa.pub을 복사해서 github settings의 deploy key에 등록하기 

# 등록후 kubectl secret 파일을 생성 필요 
# secret을 통해 생성된 경로에 file 로드하기 
kubectl create secret generic airflow-git-ssh-secret \
  --from-file=gitSshKey=/Users/lines/sources/03_onlines_corp/settings/airflow_settings/id_rsa \
  --from-file=id_ed25519.pub=/Users/lines/sources/03_onlines_corp/settings/airflow_settings/id_rsa.pub \
  -n airflow


# 이후 helm을 통해시 values.yaml 을 추출하기 
helm show values apache-airflow/airflow > values.yaml

# values.yaml 수정하기 
vi values.yaml 

  gitSync:
    enabled: true

    # git repo clone url
    # ssh example: git@github.com:apache/airflow.git
    # https example: https://github.com/apache/airflow.git
    repo: ${Organization-ID}@github.com:${Orgarnizations}/${Project-Repo-Name}.git
    branch: master
    rev: HEAD
    # The git revision (branch, tag, or hash) to check out, v4 only
    ref: master
    depth: 1
    # the number of consecutive failures allowed before aborting
    maxFailures: 0
    # subpath within the repo where dags are located
    # should be "" if dags are at repo root
    subPath: ""
    # if your repo needs a user name password
    # you can load them to a k8s secret like the one below
    #   ---
    #   apiVersion: v1
    #   kind: Secret
    #   metadata:
    #     name: git-credentials
    #   data:
    #     # For git-sync v3
    #     GIT_SYNC_USERNAME: <base64_encoded_git_username>
    #     GIT_SYNC_PASSWORD: <base64_encoded_git_password>
    #     # For git-sync v4
    #     GITSYNC_USERNAME: <base64_encoded_git_username>
    #     GITSYNC_PASSWORD: <base64_encoded_git_password>
    # and specify the name of the secret below
    #
    # credentialsSecret: git-credentials
    #
    #
    # If you are using an ssh clone url, you can load
    # the ssh private key to a k8s secret like the one below
    #   ---
    #   apiVersion: v1
    #   kind: Secret
    #   metadata:
    #     name: airflow-ssh-secret
    #   data:
    #     # key needs to be gitSshKey
    #     gitSshKey: <base64_encoded_data>
    # and specify the name of the secret below
    sshKeySecret: airflow-git-ssh-secret
    
# 수정후 helm 재 반영하기 
helm upgrade --install airflow apache-airflow/airflow -n airflow -f values.yaml --debug

# 외부 노출 시키기 
kubectl expose deployment airflow-webserver --type=LoadBalancer --name=airflow-service
 
# 외부 노출 서비스 체크하기 
kubectl get services
```

# AirFlow Study

##  Get Started 

- [Get Started](https://airflow.apache.org/docs/apache-airflow/stable/start.html)

기본적인 사용 용례에 대한 정리   

airflow는 python 기반으로 되어 있으며, all-in-one standalone 버전으로 아래의 명령어로 사용이 가능하다. 

```shell 
airflow standalone 
```

all in one 버전으로 사용하지 않고 개별적으로 사용하고 싶을 경우 아래와 같이 명령어를 분리해서 사용이 가능하다. 

```shell 
airflow db migrate

airflow users create \
    --username admin \
    --firstname Peter \
    --lastname Parker \
    --role Admin \
    --email spiderman@superhero.org

airflow webserver --port 8080

airflow scheduler
```

서비스 기동후 Dags 중에 테스트 Dags를 정의해서 아래와 같이 명령어로 실행할 수 있다. 

```shell 
# run your first task instance
airflow tasks test example_bash_operator runme_0 2015-01-01
# run a backfill over 2 days
airflow dags backfill example_bash_operator \
    --start-date 2015-01-01 \
    --end-date 2015-01-02
```

위의 명령어 실행시 동작하는 데이터는 본인이 사용하는 bash 또는 zsh의 환경변수 경로로 지정한 폴더 내에 로그 및 데이터가 생성된다. 

```shell 
# ~/.zshrc 파일 내 추가된 경로 / 해당 폴더경로는 생성되어 있어야함! 

export AIRFLOW_HOME=~/airflow
```

- [Airflow를 활용할 수 있는 다양한 방법들](https://airflow.apache.org/docs/apache-airflow/stable/installation/index.html#using-production-docker-images)

- [Airflow 필요 환경 및 조건 정의](https://airflow.apache.org/docs/apache-airflow/stable/installation/prerequisites.html)

- [Airflow CLI](https://airflow.apache.org/docs/apache-airflow/stable/howto/usage-cli.html)

Airflow 자동완성 적용하기 Profile에 추가해서 사용하기 

```shell 
register-python-argcomplete airflow >> ~/.bashrc
```

Airflow의 Dags를 추출하는 방법에 대해서 알아보면, 

```shell 
airflow dags show example_complex
```

```shell 
airflow dags show example_bash_operator --save example_bash_operator.png

airflow dags show example_bash_operator --imgcat
```

**Command Formatting Output**     

[Formatting commands output](https://airflow.apache.org/docs/apache-airflow/stable/howto/usage-cli.html#formatting-commands-output).  


**Purge history from metadata database**     

- [https://airflow.apache.org/docs/apache-airflow/stable/howto/usage-cli.html#purge-history-from-metadata-database](https://airflow.apache.org/docs/apache-airflow/stable/howto/usage-cli.html#purge-history-from-metadata-database)

### install python airflow 

- [airflow install](https://airflow.apache.org/docs/apache-airflow/stable/installation/installing-from-pypi.html)

```
pip install "apache-airflow==2.8.2" apache-airflow-providers-google==10.1.0
```

## AirFlow Documentation Help

### AirFlow Docker Image 

- [Docker Image for Apache Airflow](https://airflow.apache.org/docs/docker-stack/index.html)

### Airflow on Kubernetes with Heml Chart  

- [kubernetes With Airflow](https://airflow.apache.org/docs/helm-chart/stable/index.html)

### Prerequisites Of Airflow 

- [Airflow - Prerequisites](https://airflow.apache.org/docs/apache-airflow/stable/installation/prerequisites.html#prerequisites)

### LifeCycle and System EOL for Airflow  

- [Airflow System EOL](https://airflow.apache.org/docs/apache-airflow/stable/installation/prerequisites.html#prerequisites)

### Upgrade and migrate for Airflow 

- [Airflow System Migrations](https://airflow.apache.org/docs/apache-airflow/stable/installation/upgrading.html)

### Security for Airflow 

- [Airflow Security 구성](https://airflow.apache.org/docs/apache-airflow/stable/security/index.html)

### Setting of AirFlow RateLimiting

- [Airflow Rate Limiting](https://airflow.apache.org/docs/apache-airflow/stable/security/webserver.html#rate-limiting)

### Handling Masking on Sensitive Information 

- [민감한 정보에 대한 Masking 처리](https://airflow.apache.org/docs/apache-airflow/stable/security/webserver.html#rate-limiting)

# AirFlow Basic

## Architecture Overview 

### Control Flow 

#### XCom 

- [Airflow XCom](https://dydwnsekd.tistory.com/107)

XCom은 Dag 내의 task 사이에서 데이터를 전달하기 위해서 사용되는데, Celery Executor를 예로 들면, 각 Task들이 각기 다른 Woker에서 실행쇨 수 있으면,  
XCom은 이러한 경우 task간 데이터 전달을 가능하게 한다. 

- Xcom 사용하기 
  - Python Operator return 값을 통한 Xcom 사용 
  - push-pull 을 이용한 Xcom 사용 
  - Jinja Template 을 이용한 Xcom 사용 

#### PythonOperator return 값을 이용한 Xcom 사용 

```python
def return_xcom():
    return "xcom!"
    
return_xcom = PythonOperator(
    task_id = 'return_xcom',
    python_callable = return_xcom,
    dag = dag
)
```

#### push-pull 을 이용한 Xcom 사용 

PythonOperator에서 return을 하는 방법 이외에도 아래와 같이 context['task_instance']를 이용하여 xcom에 push, pull 하여 데이터를 주고받는 것 또한 가능한데, 여기서 알아야 할 내용들에 대해 간단하게 설명한다.
먼저 context['task_instance']와 context['ti']는 동일한 의미로 ti = task_instance로 간단하게 축약하여 사용할 수 있다.
다음으로 PythonOperator을 사용하는 경우 return과 push를 하나의 task에서 중복하여 사용할 수 있으며, 해당 데이터를 전달받는 곳에서 xcom_pull(key=~) 혹은 xcom_pull(task_ids=~)를 이용해 전달받는 방식이 서로 다른 것을 인지하자.
return으로 xcom을 사용하는 경우 xcom_pull(task_ids)를 사용해 데이터를 전달받고,
push하는 경우에는 key-value 형식에 따라 데이터를 주고받게 된다

```python
def xcom_push_test(**context):
    xcom_value = "xcom_push_value"
    context['task_instance'].xcom_push(key='xcom_push_value', value=xcom_value)

    return "xcom_return_value"

def xcom_pull_test(**context):
    xcom_return = context["task_instance"].xcom_pull(task_ids='return_xcom')
    xcom_push_value = context['ti'].xcom_pull(key='xcom_push_value')
    xcom_push_return_value = context['ti'].xcom_pull(task_ids='xcom_push_task')

    print("xcom_return : {}".format(xcom_return))
    print("xcom_push_value : {}".format(xcom_push_value))
    print("xcom_push_return_value : {}".format(xcom_push_return_value))
    
xcom_push_task = PythonOperator(
    task_id = 'xcom_push_task',
    python_callable = xcom_push_test,
    dag = dag
)

xcom_pull_task = PythonOperator(
    task_id = 'xcom_pull_task',
    python_callable = xcom_pull_test,
    dag = dag
)
```

#### Jinja templates를 이용한 Xcom 사용 

```python
bash_xcom_taskids = BashOperator(
    task_id='bash_xcom_taskids',
    bash_command='echo "{{ task_instance.xcom_pull(task_ids="xcom_push_task") }}"',
    dag=dag
)

bash_xcom_key = BashOperator(
    task_id='bash_xcom_key',
    bash_command='echo "{{ ti.xcom_pull(key="xcom_push_value") }}"',
    dag=dag
)

bash_xcom_push = BashOperator(
    task_id='bash_xcom_push',
    bash_command='echo "{{ ti.xcom_push(key="bash_xcom_push", value="bash_xcom_push_value") }}"',
    dag=dag
)

bash_xcom_pull = BashOperator(
    task_id='bash_xcom_pull',
    bash_command='echo "{{ ti.xcom_pull(key="bash_xcom_push") }}"',
    dag=dag
)
```

#### 전체 예제 ( Controll Flow with XCom )

```python
from airflow import DAG
from airflow.operators.python import PythonOperator
from airflow.operators.bash import BashOperator
from datetime import datetime

dag = DAG(
    dag_id = 'xcom_test',
    start_date = datetime(2021,9,23),
    catchup=False,
    schedule_interval='@once'
)

def return_xcom():
    return "xcom!"

def xcom_push_test(**context):
    xcom_value = "xcom_push_value"
    context['task_instance'].xcom_push(key='xcom_push_value', value=xcom_value)

    return "xcom_return_value"

def xcom_pull_test(**context):
    xcom_return = context["task_instance"].xcom_pull(task_ids='return_xcom')
    xcom_push_value = context['ti'].xcom_pull(key='xcom_push_value')
    xcom_push_return_value = context['ti'].xcom_pull(task_ids='xcom_push_task')

    print("xcom_return : {}".format(xcom_return))
    print("xcom_push_value : {}".format(xcom_push_value))
    print("xcom_push_return_value : {}".format(xcom_push_return_value))


return_xcom = PythonOperator(
    task_id = 'return_xcom',
    python_callable = return_xcom,
    dag = dag
)

xcom_push_task = PythonOperator(
    task_id = 'xcom_push_task',
    python_callable = xcom_push_test,
    dag = dag
)

xcom_pull_task = PythonOperator(
    task_id = 'xcom_pull_task',
    python_callable = xcom_pull_test,
    dag = dag
)

bash_xcom_taskids = BashOperator(
    task_id='bash_xcom_taskids',
    bash_command='echo "{{ task_instance.xcom_pull(task_ids="xcom_push_task") }}"',
    dag=dag
)

bash_xcom_key = BashOperator(
    task_id='bash_xcom_key',
    bash_command='echo "{{ ti.xcom_pull(key="xcom_push_value") }}"',
    dag=dag
)

bash_xcom_push = BashOperator(
    task_id='bash_xcom_push',
    bash_command='echo "{{ ti.xcom_push(key="bash_xcom_push", value="bash_xcom_push_value") }}"',
    dag=dag
)

bash_xcom_pull = BashOperator(
    task_id='bash_xcom_pull',
    bash_command='echo "{{ ti.xcom_pull(key="bash_xcom_push") }}"',
    dag=dag
)

return_xcom >> xcom_push_task >>xcom_pull_task >> bash_xcom_taskids >> bash_xcom_key >> bash_xcom_push >> bash_xcom_pull
```


## [Core Concepts > Dags](https://airflow.apache.org/docs/apache-airflow/stable/core-concepts/dags.html)

### Dag Definition


#### Task Dependencies 

- >>, << 를 이용한 흐름 구성 가능 

```python
with DAG(
        dag_id=dag_name,
        default_args=default_args,
        description='Database Schema Sync',
        schedule=timedelta(days=1),
        start_date=datetime(2023, 1, 1, 17, 30, tzinfo=seoul_time),
        catchup=False,
        tags=['etl']
) as dag:
    task1 = EmptyOperator(task_id="task1")

    task2 = EmptyOperator(task_id="task2")

    task3 = EmptyOperator(task_id="task3")

    task4 = EmptyOperator(task_id="task4")

    task1 >> [task2, task3]
    task1 << task4
```


- upstream, downstream 구성을 이용한 workflow 설계 


```python
with DAG(
        dag_id=dag_name,
        default_args=default_args,
        description='Database Schema Sync',
        schedule=timedelta(days=1),
        start_date=datetime(2023, 1, 1, 17, 30, tzinfo=seoul_time),
        catchup=False,
        tags=['etl']
) as dag:
    task1 = EmptyOperator(task_id="task1")

    task2 = EmptyOperator(task_id="task2")

    task3 = EmptyOperator(task_id="task3")

    task4 = EmptyOperator(task_id="task4")

    task1.set_downstream([task2, task3])
    task3.set_upstream(task4)
```

#### Dag Loading
#### Dag Execution

```python 
with DAG("my_daily_dag", schedule="0 0 * * *"):
    ...

with DAG("my_one_time_dag", schedule="@once"):
    ...

with DAG("my_continuous_dag", schedule="@continuous"):
    ...
```

#### Dag Assign

- with DAG
- @dag 
- Upstream, Downstream 

#### Default Args 

하나의 Dag 안에 많은 오퍼레이터가 존재하는 경우, 기본 인자를 설정할 필요가 있다. default_args를 사용할 경우 해당 Dag에 연결되어 있는 Operator는 defualt_args에 정의된 값은 모두 같이 사용한다. 


```python
import pendulum

with DAG(
    dag_id="my_dag",
    start_date=pendulum.datetime(2016, 1, 1),
    schedule="@daily",
    default_args={"retries": 2},
):
    op = BashOperator(task_id="hello_world", bash_command="Hello World!")
    print(op.retries)  # 2
```
#### Dag Decorator 

Dag를 정의하는 방식은 

- DAG()를 사용하는 방법이 있고, 
- @dag를 이용하는 방법이 있고, 

```python 
@dag(
    schedule=None,
    start_date=pendulum.datetime(2021, 1, 1, tz="UTC"),
    catchup=False,
    tags=["example"],
)
def example_dag_decorator(email: str = "example@example.com"):
    """
    DAG to send server IP to email.

    :param email: Email to send IP to. Defaults to example@example.com.
    """
    get_ip = GetRequestOperator(task_id="get_ip", url="http://httpbin.org/get")

    @task(multiple_outputs=True)
    def prepare_email(raw_json: dict[str, Any]) -> dict[str, str]:
        external_ip = raw_json["origin"]
        return {
            "subject": f"Server connected from {external_ip}",
            "body": f"Seems like today your server executing Airflow is connected from IP {external_ip}<br>",
        }

    email_info = prepare_email(get_ip.output)

    EmailOperator(
        task_id="send_email", to=email, subject=email_info["subject"], html_content=email_info["body"]
    )


example_dag = example_dag_decorator()
```

#### Control Flow  
##### Branching

```python
@task.branch(task_id="branch_task")
def branch_func(ti=None):
    # xcom_pull을 이용해서 start_task 에서 값을 전달 받아서 if.else 분기를 통해 Branching이 가능하다. 
    xcom_value = int(ti.xcom_pull(task_ids="start_task"))
    if xcom_value >= 5:
        return "continue_task"
    elif xcom_value >= 3:
        return "stop_task"
    else:
        return None


start_op = BashOperator(
    task_id="start_task",
    bash_command="echo 5",
    do_xcom_push=True,
    dag=dag,
)

branch_op = branch_func()

continue_op = EmptyOperator(task_id="continue_task", dag=dag)
stop_op = EmptyOperator(task_id="stop_task", dag=dag)

start_op >> branch_op >> [continue_op, stop_op]
```


##### Latest Only 

```python 
import datetime

import pendulum

from airflow.models.dag import DAG
from airflow.operators.empty import EmptyOperator
from airflow.operators.latest_only import LatestOnlyOperator
from airflow.utils.trigger_rule import TriggerRule

with DAG(
    dag_id="latest_only_with_trigger",
    schedule=datetime.timedelta(hours=4),
    start_date=pendulum.datetime(2021, 1, 1, tz="UTC"),
    catchup=False,
    tags=["example3"],
) as dag:
    latest_only = LatestOnlyOperator(task_id="latest_only")
    task1 = EmptyOperator(task_id="task1")
    task2 = EmptyOperator(task_id="task2")
    task3 = EmptyOperator(task_id="task3")
    task4 = EmptyOperator(task_id="task4", trigger_rule=TriggerRule.ALL_DONE)

    latest_only >> task1 >> [task3, task4]
    task2 >> [task3, task4]
```

#### Depends On Past 


##### Trigger Rules

- Branch without Trigger

```python
# dags/branch_without_trigger.py
import pendulum

from airflow.decorators import task
from airflow.models import DAG
from airflow.operators.empty import EmptyOperator

dag = DAG(
    dag_id="branch_without_trigger",
    schedule="@once",
    start_date=pendulum.datetime(2019, 2, 28, tz="UTC"),
)

run_this_first = EmptyOperator(task_id="run_this_first", dag=dag)


@task.branch(task_id="branching")
def do_branching():
    return "branch_a"


branching = do_branching()

branch_a = EmptyOperator(task_id="branch_a", dag=dag)
follow_branch_a = EmptyOperator(task_id="follow_branch_a", dag=dag)

branch_false = EmptyOperator(task_id="branch_false", dag=dag)

join = EmptyOperator(task_id="join", dag=dag)

run_this_first >> branching
branching >> branch_a >> follow_branch_a >> join
branching >> branch_false >> join
```

- With Trigger 

Trigger Rule로 정의 하면, 

- all_success 
- all_failed 
- all_done 
- all_skipped 
- one_failed 
- one_success 
- one_done 
- none_failed 
- none_failed_min_one_success 
- none_skipped 
- always 

```python
import datetime

import pendulum

from airflow.models.dag import DAG
from airflow.operators.empty import EmptyOperator
from airflow.operators.latest_only import LatestOnlyOperator
from airflow.utils.trigger_rule import TriggerRule

with DAG(
    dag_id="latest_only_with_trigger",
    schedule=datetime.timedelta(hours=4),
    start_date=pendulum.datetime(2021, 1, 1, tz="UTC"),
    catchup=False,
    tags=["example3"],
) as dag:
    latest_only = LatestOnlyOperator(task_id="latest_only")
    task1 = EmptyOperator(task_id="task1")
    task2 = EmptyOperator(task_id="task2")
    task3 = EmptyOperator(task_id="task3")
    task4 = EmptyOperator(task_id="task4", trigger_rule=TriggerRule.ALL_DONE)

    latest_only >> task1 >> [task3, task4]
    task2 >> [task3, task4]
```


#### Setup and teardown 

- [Setup And Teardown](https://airflow.apache.org/docs/apache-airflow/stable/howto/setup-and-teardown.html)

#### Dynamics DAGs 

for loop 문장 등을 이용해서 동적으로 dag 내의 Task를 생성 및 연결할 수 있다. 

```python
 with DAG("loop_example", ...):
     first = EmptyOperator(task_id="first")
     last = EmptyOperator(task_id="last")

     options = ["branch_a", "branch_b", "branch_c", "branch_d"]
     for option in options:
         t = EmptyOperator(task_id=option)
         first >> t >> last
```

#### DAG Visualization 

##### TaskGroups

Dag Graph 에서 Task를 Group 단위로 묶어서 표현할 수 있다. 

```python
 from airflow.decorators import task_group


 @task_group()
 def group1():
     task1 = EmptyOperator(task_id="task1")
     task2 = EmptyOperator(task_id="task2")


 task3 = EmptyOperator(task_id="task3")

 group1() >> task3
```

```python
import datetime

from airflow import DAG
from airflow.decorators import task_group
from airflow.operators.bash import BashOperator
from airflow.operators.empty import EmptyOperator

with DAG(
    dag_id="dag1",
    start_date=datetime.datetime(2016, 1, 1),
    schedule="@daily",
    default_args={"retries": 1},
):

    @task_group(default_args={"retries": 3})
    def group1():
        """This docstring will become the tooltip for the TaskGroup."""
        task1 = EmptyOperator(task_id="task1")
        task2 = BashOperator(task_id="task2", bash_command="echo Hello World!", retries=2)
        print(task1.retries)  # 3
        print(task2.retries)  # 2
```

##### Edge Labels 

Task 사이의 Labeling을 표현할 수 있다. 

```python

with DAG(
    "example_branch_labels",
    schedule="@daily",
    start_date=pendulum.datetime(2021, 1, 1, tz="UTC"),
    catchup=False,
) as dag:
    ingest = EmptyOperator(task_id="ingest")
    analyse = EmptyOperator(task_id="analyze")
    check = EmptyOperator(task_id="check_integrity")
    describe = EmptyOperator(task_id="describe_integrity")
    error = EmptyOperator(task_id="email_error")
    save = EmptyOperator(task_id="save")
    report = EmptyOperator(task_id="report")

    ingest >> analyse >> check
    check >> Label("No errors") >> save >> report
    check >> Label("Errors found") >> describe >> error >> report
```

#### DAG & Task Documentation

- doc
- doc_json 
- doc_yaml
- doc_md 
- doc_rst 

```python 
"""
### My great DAG
"""
import pendulum

dag = DAG(
    "my_dag",
    start_date=pendulum.datetime(2021, 1, 1, tz="UTC"),
    schedule="@daily",
    catchup=False,
)
dag.doc_md = __doc__

t = BashOperator("foo", dag=dag)
t.doc_md = """\
#Title"
Here's a [url](www.airbnb.com)
"""
```

#### SubDAGs

#### TaskGroups vs SubDAGs

#### Packaging DAGs

### DAG Rungs

### Operators

- [Operators and Hooks Reference](https://airflow.apache.org/docs/apache-airflow/stable/operators-and-hooks-ref.html)

#### Operator Package 

- [Airflow.sensors](https://airflow.apache.org/docs/apache-airflow/2.2.3/_api/airflow/sensors/index.html)

#### KubernetesPodOnOperator 

- [Kubernetes Pod On Operator](https://airflow.apache.org/docs/apache-airflow-providers-cncf-kubernetes/stable/_modules/index.html)

```python
spark_submit_sample = KubernetesPodOperator(
    task_id='spark_submit_sample',
    name='spark_submit_sample',
    namespace='airflow',
    image='spark_client_1:1.0',
    arguments=["spark", SparkSubmitCommandTool.getCommand(SparkSubmitCommandTool(),
                        class_name='com.linecorp.spark.application.SampleApplication',
                        master='yarn',
                        keytab='user.keytab',
                        principal='user@FINANCIAL.HADOOP.DATA.COM',
                        deploy_mode='cluster',
                        driver_cores='1',
                        driver_memory='2g',
                        executor_cores='2',
                        executor_memory='4',
                        num_executors='5',
                        jars='hdfs://line_financial/spark_app/spark-sample.jar',
                        args=[execution_date],
                        conf=["spark.dynamicAllocation.enabled=true",
                              "spark.shuffle.service.enabled=true",
                              "spark.dynamicAllocation.minExecutors=5",
                              "spark.dynamicAllocation.maxExecutors=10",
                              "spark.rpc.message.maxSize=256"])],
    resources={ 'request_cpu': '1000m',
                'request_memory': '2048Mi',
                'limit_cpu': '2000m',
                'limit_memory': '4095Mi'},
    hostnetwork=True,
    in_cluster=True,
    is_delete_operator_pod=True,
    startup_timeout_seconds=180,
    execution_timeout=timedelta(minutes=120),
    retries=2,
    retry_delay=timedelta(minutes=2),
    on_retry_callback=SlackTool.make_handler(slack_channel,
                                             color='warning',
                                             message="Retry Task"),
    image_pull_policy='IfNotPresent',
    service_account_name='airflow',
    dag=dag)
```

#### Kubernetes Resource Operator 

- [Kubernetes Resource Operator](https://airflow.apache.org/docs/apache-airflow-providers-cncf-kubernetes/stable/_modules/tests/system/providers/cncf/kubernetes/example_kubernetes_resource.html)

- kubernetes namespace 생성을 위한 권한이 필요함.

```yaml 
# Create a cluster role that allowed to perform
# ["get", "list", "create", "delete", "patch"] over ["namespaces"]
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: airflow-worker
rules:
  - apiGroups: [""]
    resources: ["namespaces"]
    verbs: ["get", "list", "create", "delete", "patch"]
---
# Associate the cluster role with the service account
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: airflow-worker
  # make sure NOT to mention 'namespace' here or
  # the permissions will only have effect in the
  # given namespace
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: airflow-worker
subjects:
- kind: ServiceAccount
  name: airflow-worker
  namespace: airflow
```

```shell 
kubectl apply -f rbac.yaml
```

- pods에 대한 권한도 추가로 적용하기 

```yaml 
# Create a cluster role that allowed to perform
# ["get", "list", "create", "delete", "patch"] over ["namespaces"]
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: airflow-worker
rules:
  - apiGroups: [""]
    resources: ["namespaces", "pods"]
    verbs: ["get", "list", "create", "delete", "patch"]
---
# Associate the cluster role with the service account
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: airflow-worker
  # make sure NOT to mention 'namespace' here or
  # the permissions will only have effect in the
  # given namespace
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: airflow-worker
subjects:
- kind: ServiceAccount
  name: airflow-worker
  namespace: airflow
```

- kubectl on airflow 

  - https://stackoverflow.com/questions/68195000/how-to-get-the-kubernetes-api-key

문제 해결 필요 사항 - 쿠버네티스 API를 airflow worker에서 호출할 때 api 호출에 대한 외부 허용에 대한 rule을 정리해봐야함. namespace가 다른 kube-system으로 호출하는 방법은?....?   
이건 airflow의 문제가 아닌 api 호출 구조에 대한 정리하는 것이 중요한 포인트로 보임. 아하 이제 이해함...!


# 운영 관점에서 알고 있어야할 링크

- [Logging & Monitoring](https://airflow.apache.org/docs/apache-airflow/stable/administration-and-deployment/logging-monitoring/index.html)
  - [Health Check](https://airflow.apache.org/docs/apache-airflow/stable/administration-and-deployment/logging-monitoring/check-health.html)
  - [Database Backend](https://airflow.apache.org/docs/apache-airflow/stable/howto/set-up-database.html#setting-up-a-postgresql-database)
  - [Metrics](https://airflow.apache.org/docs/apache-airflow/stable/administration-and-deployment/logging-monitoring/metrics.html)

# AirFlow 관련 자료

- [오늘의 집 - Airflow](https://www.bucketplace.com/post/2021-04-13-%EB%B2%84%ED%82%B7%ED%94%8C%EB%A0%88%EC%9D%B4%EC%8A%A4-airflow-%EB%8F%84%EC%9E%85%EA%B8%B0/)
- [엄청 자세한 튜토리얼 - 왕초보자용](https://velog.io/@clueless_coder/Airflow-%EC%97%84%EC%B2%AD-%EC%9E%90%EC%84%B8%ED%95%9C-%ED%8A%9C%ED%86%A0%EB%A6%AC%EC%96%BC-%EC%99%95%EC%B4%88%EC%8B%AC%EC%9E%90%EC%9A%A9)
-[케르베로스(Kerberos)란? 동작이해하기](https://juhi.tistory.com/75)
- [SBOM(Software Bill Of Materials)](https://openbee.kr/444)
- [Dags 등록 및 사용하기](https://velog.io/@inhwa1025/Airflow-Web-UI-%EC%82%AC%EC%9A%A9%ED%95%98%EA%B8%B0-%EB%B0%8F-DAG-%EB%93%B1%EB%A1%9D%ED%95%98%EA%B8%B0)
- [Airflow Operator - Tistory 50](https://mightytedkim.tistory.com/50)
- [Airflow Operator - Tistory 57](https://mightytedkim.tistory.com/57)
- [Airflow on Kubernetes: Get Started in 10 mins](https://marclamberti.com/blog/airflow-on-kubernetes-get-started-in-10-mins/)
- [로컬 환경에서 쿠버네티스 구축하기 1](https://pearlluck.tistory.com/794)
- [로컬 환경에서 쿠버네티스 구축하기 2](https://pearlluck.tistory.com/795)
- [Production Guide](https://airflow.apache.org/docs/helm-chart/stable/production-guide.html#accessing-the-airflow-ui)
- [Airflow Kubernetes 처리](https://airflow.apache.org/docs/apache-airflow-providers-cncf-kubernetes/stable/connections/kubernetes.html)

# Airflow 훌륭한 레퍼런스 

- [Build A Data Warehouse using MYSQL, Airflow and DBT](https://github.com/nebasam/Data-Warehouse-Tech-Stack/tree/main)

# Airflow Sensor 활용하기 

- [AirFlow의 Sensor 이해하기](https://www.bearpooh.com/153)
    - [SQLSensor 사용 방법](https://mvje.tistory.com/191)
    - [SQL의 데이터 정보 변경시 데이터를 추출하여 데이터 적재 및 갱신 프로세스 처리](https://medium.com/@kausarbazla/airflow-sensors-operators-hooks-99cb077c5dd1)
    - [Airflow Sensor King](https://marclamberti.com/blog/airflow-sensors/)