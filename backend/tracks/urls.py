from django.urls import path

from .views import (
    TrackListAPIView,
    TrackGetUpdateAPIView,
    TrackCoverGetAPIView,
    TrackQueueAPIView,
)

urlpatterns = [
    path("", TrackListAPIView.as_view(), name="tracks-list"),
    path("queue/", TrackQueueAPIView.as_view(), name="tracks-queue"),
    path("<int:pk>/", TrackGetUpdateAPIView.as_view()),
    path("cover/<int:pk>/", TrackCoverGetAPIView.as_view(), name="tracks-cover"),
]
