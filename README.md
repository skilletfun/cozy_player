# Build container

```
podman build -t cozy-player .
```

# Run container

```
podman run -d -p 8000:80 \
-v cozy-player-db:/db \
-v ~/Develop/cozy_player/.music:/music \
--name cozy-player --replace cozy-player:latest
```

# Push container to Docker Hub

```
podman push <container-id> docker://docker.io/skilletfun/cozy-player:<tag>
```

# TODO:

- переехать на mongo / postgres (или хотя бы добавить возможность выбора)
- в списке треков должно быть разделение по альбомам / синглам / EP
- playlists
