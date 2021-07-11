# Contributing

## Main branch is always deployable

To support continuous delivery the branch called `main` is always deployable.
Issues are always created off of the `main` branch.

Issue branches are always merged into the `main` branch. See [Git Pull Request Process](#Git-Pull-Request Process-every-issue-is-a-branch).

## Git Pull Request Process- every issue is a branch

> In short: Always create an issue first, get it's issue number and prefix your branch name with
  the issue number. See also [Everything starts with an issue](https://about.gitlab.com/blog/2016/03/03/start-with-an-issue/#:~:text=%E2%80%9CAlways%20start%20with%20an%20issue,but%20the%20impact%20is%20huge.)

Question:
> I would like to ask for an explanation about git best practices. Like naming branches etc, I was in most cases working alone on projects so things like this were not important to me but now I need some guidance about these things. 
Answer: 

- Contributions and changes must start with creating an issue, even if a tiny change. This gives you an issue number to use later on, the issue number is the branch name prefix. E.g. issue `123-<human-readable-name>` this way, you always know
which issue a branch relates to.


## How to create issue branches

Locally:

- Checkout to main branch `git checkout main` (because you always want to branch of the most up to date code)
- Make sure your main branch is up to date.
  - `git fetch`  (this fetches all changes into the .git folder , but it does not change your files yet)
  - `git rebase origin/main` (this does change your files and makes them up to date with main) and is safe to do- never rebase a shared branch you're collaborating on

- Create a new branch for the issue, using the number
  - e.g. `git checkout 376-add-edit-product`
  - (remember you're on the `main branch`, which is good because it contains the most up to date code because you fetched, then rebased) 
- `git checkout -b <issue-number>-name-of-issue` 
- do your coding , creating small commits which reference the issue e.g. if you updated template.html file
- `git add <file>` (or which ever files you've added/changed)
- `git commit -m "Fix #376 add edit product"`
- when you use the "#<issue-number>" in a commit message, GitHub automatically shows that commit on the GitHub issue. This is very useful for seeing the issue/code relationship.

- When you say `git commit -m "Fix #<issue number> my comment about the code`, if the commit is merged into main, then it automatically close the issue (great time saver!)
- If you're finished , push the branch , but wait!
  - Maybe the main branch has more changes since you wrote your code... you need to fetch and apply those first
- `git fetch origin/main` (fetches any changes which happened whilst you were working)
- `git pull origin/main`n  (applies those changes to your current branch) notice we're using `pull` and not `rebase` in case we're collaborating with another person
- Now you're ready to push! Finally!
- `git push origin 376-<issue-name>`
- Never use force push if your working with others in the branch, it will destroy their work. If you're not 100% sure, don't ever use force push. It's rarely needed
- Go to GitHub and raise a pull request 
- There is no need to delete the branch remotely because branches are copy on write (very tiny file size) but locally you might want to delete them if you have hundreds or don't want to see them anymore (`git branch -d <branch-name>`)
- It's normally unhelpful to delete a local branch 10mins after raising a pull request, because you might need to add things later on. You can always get the branch back, it takes seconds, but it's a hassle. Just keep the branch locally , it's not causing any problems, it's helping you because you might want to go back to it in a few weeks.
