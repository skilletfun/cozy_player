from django.urls import path

from .views import PlaylistListCreateAPIView, PlaylistRUDAPIView


urlpatterns = [
    path("", PlaylistListCreateAPIView.as_view(), name="playlists-list"),
    path("<int:pk>/", PlaylistRUDAPIView.as_view(), name="playlists-rud"),
]
