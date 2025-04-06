from django.urls import path

from .views import ArtistListAPIView, ArtistGetAPIView, ArtistCoverGetAPIView

urlpatterns = [
    path("", ArtistListAPIView.as_view()),
    path("<int:pk>/", ArtistGetAPIView.as_view()),
    path("cover/<int:pk>/", ArtistCoverGetAPIView.as_view()),
]
