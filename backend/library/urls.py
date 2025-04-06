from django.urls import path

from .views import RescanLibraryAPIView


urlpatterns = [
    path("rescan/", RescanLibraryAPIView.as_view()),
]
