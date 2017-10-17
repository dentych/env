# Git initial setup
if [[ $SETUPECHO = true ]]; then
    echo "Setting up git username and email..."
fi
git config --global user.name "$GITUSERNAME"
git config --global user.email "$GITEMAIL"

# Git stuff
if [[ $SETUPECHO = true ]]; then
    echo "Setting up git aliases..."
fi
git config --global alias.cp "cherry-pick"
git config --global alias.co "checkout"
git config --global alias.cl "clone"
git config --global alias.c "commit"
git config --global alias.st "status -sb"
git config --global alias.br "branch"
git config --global alias.d "diff"
git config --global alias.dc "diff --cached"
git config --global alias.p "pull -p"
git config --global alias.pu "push -u"
git config --global alias.f "fetch -p"
git config --global alias.b "branch"
git config --global alias.logn "log --all --graph --oneline --decorate"
git config --global alias.lognb "log --graph --oneline --decorate"
git config --global alias.pushb "!git push origin \$(git rev-parse --abbrev-ref HEAD) -u"