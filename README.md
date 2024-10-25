# Backend Technical Test at Scalingo

## Specifications

The architecture is a simplified version of a Clean Architecture under a single microservice.
I usually like to do it with a domain layer and cqrs, but there is no need for it here.

The Application is divided into three layers:

- Interface Layer: The entry point of the application, where the request is received and the response is sent. This
  layer is calling the application layer
- Service Layer: This layer is responsible for holding the use cases which are going to call the domain (which is
  empty in this case) and the adapter layer.
- Adapter Layer: Adapters are connectors to external services or databases. In this case, in this case the adapter is
  responsible for calling the GitHub API and aggregating the languages of the repositories. It could be a bit confusing
  since the internal adapters are called Repository but those have nothing to do with github.

I've decided to do the aggregation parts within the adapters since it could be entirely different for another source.
This allows to not introduce any breaking changes in the service layer.

### Further Improvements

- A cache would be highly beneficial to store either a previous request or a repository that has already been
  aggregated.
  We could, for example, set up an memory cache or a redis cache.
- Support for pagination
- Support for custom populating based on some 'required' parameters

### Scalability & Maintainability

- Since the application is stateless, it can be easily scaled horizontally.
  It can be deployed on multiple instances
  behind a load balancer.
- The multiple layers of the application make it easy to maintain, test it and add new features.
- Adapters are easily replaceable or mockable.
- There is mappings between the final response and the layers, so data can be easily transformed, added, removed without
  introducing breaking changes.

### Rate Limit

GitHub API has a rate limit of 60 requests per hour for unauthenticated requests.
And since aggregating the language from the repositories, need a additional request per repository,
the rate limit can be reached quite quickly.

However, you can set a token in the config.go file to increase the rate
limit to 5000 requests per hours.

## Execution

```
docker compose up
```

Application will be then running on port `5000`

## Usages

```
localhost:5000/v1/repositories?limit=20&language=Ruby
```

## Performance

Performance wise, it seems that GitHub API is quite slow (Between 1 seconde and 5 secondes during my
test).

Processing the repositories with their languages for 100 repos took under 0.10ms
but could be highly inaccurate due to the GitHub API response time and various another reason.

There is a env variable 'WORKERS_POOL_SIZE'
that can be adjusted to increase the parallelism of the aggregations.