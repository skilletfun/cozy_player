from django.urls import path

from .views import ArtistListAPIView, ArtistGetAPIView, ArtistCoverGetAPIView


urlpatterns = [
    path("", ArtistListAPIView.as_view(), name="artists-list"),
    path("<int:pk>/", ArtistGetAPIView.as_view(), name="artists-retrieve"),
    path("cover/<int:pk>/", ArtistCoverGetAPIView.as_view(), name="artists-cover"),
]
