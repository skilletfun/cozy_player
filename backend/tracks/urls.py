from django.urls import path

from .views import TrackListAPIView, TrackGetAPIView

urlpatterns = [
    path("", TrackListAPIView.as_view()),
    path("<int:pk>/", TrackGetAPIView.as_view()),
]
