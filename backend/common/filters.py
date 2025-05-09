import django_filters
from django.db import models


class BaseFilterSet(django_filters.FilterSet):
    class Meta:
        filter_overrides = {
            models.CharField: {
                "filter_class": django_filters.CharFilter,
                "extra": lambda f: {
                    "lookup_expr": "icontains",
                },
            },
            models.ForeignKey: {
                "filter_class": django_filters.ModelMultipleChoiceFilter,
                "extra": lambda f: {
                    "queryset": django_filters.filterset.remote_queryset(f),
                },
            },
        }
