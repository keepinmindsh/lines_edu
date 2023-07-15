# Git Command 

## Index 

- [Git Auth Login 설정](#Git-Auth-Login-설정)
- [Git Config Global User Info](#Git-Config-Global-User-Info)
- [Git Branch 원격으로 Push 하기](#Git-Branch-원격으로-Push-하기)
- [Git 되돌리기](#Git-되돌리기)

#### Git Auth Login 설정 

```shell
gh auth login

ssh-keygen -t ed25519 -C "[howard.jeong@swit.io](mailto:howard.jeong@swit.io)"

gh auth refresh -h [github.com](http://github.com/) -s admin:public_key

gh ssh-key add ~/.ssh/id_ed25519.pub

ssh-add --apple-use-keychain ~/.ssh/id_ed25519
```

#### Git Config Global User Info 

```shell
$ git config --global user.name "Your Name"
$ git config --global user.email you@example.com
```

```shell
$ git config user.name "LainyZine"
$ git config user.email lainyzine.com@gmail.com
```

#### Git Branch 원격으로 Push 하기

```shell
git push --set-upstream origin feat/howard-import-from-local
```

#### Git 되돌리기 

아래의 명령어는 지금까지 Local에서 작업한 Git Log를 출력해준다. 

```shell
git reflog {branch name}
```

특정 Commit으로 돌아가기 위해서는 아래의 명령어를 통해서 진행한다. 

```shell
# --hard 명령어는 돌아간 이후의 변경 커밋을 모두 삭제해버립니다. 
# --mixed 명령어는 변경 이력은 전부 삭제하지만 변경된 내용에 대해서는 남아있습니다.  
# --soft 명령어는 변경 이력은 전부 삭제하지만 변경된 내용에 대해서는 남아 있습니다. 
git reset --hard {commit id 또는 HEAD의 특정 순번}

git reset HEAD^
```

rebase 를 중지하고 싶다면, 

```shell
git rebase --abort 
``` 

revert 기능은 특정 커밋을 되돌리는 작업도 하나의 커밋으로 간주하여 커밋 히스토리에 추가하는 방식  
revert는 되돌릴 커밋이 중간에 있거나, 되돌린 후 어떤 커밋이 왜 revert가 되었는지 이력을 남길 수 있어서 유용  

```shell
git revert {commit id} # 특정 commit id 로 revert 처리 

git revert {commit id}..{commit id}  # 범위 지정 
```