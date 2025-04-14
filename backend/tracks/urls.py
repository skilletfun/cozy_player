from django.urls import path

from .views import TrackListAPIView, TrackGetAPIView, TrackCoverGetAPIView

urlpatterns = [
    path("", TrackListAPIView.as_view()),
    path("<int:pk>/", TrackGetAPIView.as_view()),
    path("cover/<int:pk>/", TrackCoverGetAPIView.as_view()),
]
