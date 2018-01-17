helm upgrade hello-world chart/ --set image.tag=$(git rev-list --max-count 1 HEAD)
