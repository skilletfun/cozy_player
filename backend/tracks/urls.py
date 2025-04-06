from django.urls import path

from .views import TrackListAPIView

urlpatterns = [
    path("", TrackListAPIView.as_view()),
]
