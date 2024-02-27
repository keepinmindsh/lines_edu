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

- Scheduler 
  - 예약된 워크플로우를 트리거하고, Executor에게 작업을 전달하고 실행

- Executor 
  - 실행 중인 작업을 처리하는 실행 프로그램 
  - 기본적으로 Airflow 를 설치하면 스케줄러 내부의 모든 것을 실행하지만, 
  - Production 을 위한 Airfow 작업에서는 Worker에게 작업을 전달합니다.

- DAG Directory 
  - /dags 정의한 DAG 파일을 포함하는 폴더
  - Scheduler와 Executor, 그리고 Executor가 가진 모든 Worker 가 읽는 폴더입니다.
  - Dag 내의 Task 실행 단위에서 핸들링 할 수 있는 Operator 입니다. 
    - https://airflow.apache.org/docs/apache-airflow/stable/core-concepts/operators.html

- Metadata Database 
  - scheduler, executor 및 webserver가 상태를 저장하는 데 사용하는 데이터베이스

- Webserver 
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

## AirFlow Basic

### Core Concepts > Dag

#### [Core Concepts > Dags](https://airflow.apache.org/docs/apache-airflow/stable/core-concepts/dags.html)

##### Dag Definition
##### Dag Loading
##### Dag Execution
##### Dag Assign
##### Control Flow  
###### Branching
###### Latest Only 
###### Depends On Past 
###### Trigger Rules 
##### Setup and teardown 
##### Dynamics DAGs 
##### DAG Visualization 
###### TaskGroups 
###### Edge Labels 
##### DAG & Task Documentation
##### SubDAGs
##### TaskGroups vs SubDAGs
##### Packaging DAGs

### Operators

- [Operators and Hooks Reference](https://airflow.apache.org/docs/apache-airflow/stable/operators-and-hooks-ref.html)

# AirFlow 관련 자료

- [오늘의 집 - Airflow](https://www.bucketplace.com/post/2021-04-13-%EB%B2%84%ED%82%B7%ED%94%8C%EB%A0%88%EC%9D%B4%EC%8A%A4-airflow-%EB%8F%84%EC%9E%85%EA%B8%B0/)
- [엄청 자세한 튜토리얼 - 왕초보자용](https://velog.io/@clueless_coder/Airflow-%EC%97%84%EC%B2%AD-%EC%9E%90%EC%84%B8%ED%95%9C-%ED%8A%9C%ED%86%A0%EB%A6%AC%EC%96%BC-%EC%99%95%EC%B4%88%EC%8B%AC%EC%9E%90%EC%9A%A9)
-[케르베로스(Kerberos)란? 동작이해하기](https://juhi.tistory.com/75)
- [SBOM(Software Bill Of Materials)](https://openbee.kr/444)
- [Dags 등록 및 사용하기](https://velog.io/@inhwa1025/Airflow-Web-UI-%EC%82%AC%EC%9A%A9%ED%95%98%EA%B8%B0-%EB%B0%8F-DAG-%EB%93%B1%EB%A1%9D%ED%95%98%EA%B8%B0)

# AirFlow 관련 자료 

- [Airflow Operator - Tistory 50](https://mightytedkim.tistory.com/50)
- [Airflow Operator - Tistory 57](https://mightytedkim.tistory.com/57)
