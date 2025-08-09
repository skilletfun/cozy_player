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

# TODO:

- rescan должен влиять на очередь воспроизведения, там могут быть невалидные треки

