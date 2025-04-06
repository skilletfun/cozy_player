from django.urls import include, path


urlpatterns = [
    path("api/artists/", include("artists.urls")),
    path("api/tracks/", include("tracks.urls")),
    path("api/library/", include("library.urls")),
]
