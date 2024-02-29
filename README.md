git init

git add README.md

git commit -m "first commit"

git remote add origin git@github.com:appleconda/q2

git push -u origin master

git branch dev

git switch dev

touch src.go

git add src.go

git commit

git push origin dev

git switch master

git merge --squash dev

git add .

git commit -m "merge dev"

git push origin master
