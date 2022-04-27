
cat -p manifest.yml | yq  '(.applications.[] | select(.name == "mcs-devops-2")).version = v0.13.0'