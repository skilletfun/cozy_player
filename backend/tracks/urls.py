from django.urls import path

from .views import TrackListAPIView, TrackGetUpdateAPIView, TrackCoverGetAPIView

urlpatterns = [
    path("", TrackListAPIView.as_view()),
    path("<int:pk>/", TrackGetUpdateAPIView.as_view()),
    path("cover/<int:pk>/", TrackCoverGetAPIView.as_view()),
]
