from django.urls import path

from .views import (
    TrackListAPIView,
    TrackGetUpdateAPIView,
    TrackCoverGetAPIView,
    TrackQueueAPIView,
)

urlpatterns = [
    path("", TrackListAPIView.as_view()),
    path("queue/", TrackQueueAPIView.as_view()),
    path("<int:pk>/", TrackGetUpdateAPIView.as_view()),
    path("cover/<int:pk>/", TrackCoverGetAPIView.as_view()),
]
