from django.urls import path

from .views import RescanLibraryAPIView, StatisticAPIView


urlpatterns = [
    path("stats/", StatisticAPIView.as_view()),
    path("rescan/", RescanLibraryAPIView.as_view()),
]
