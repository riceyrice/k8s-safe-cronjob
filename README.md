# k8s-safe-cronjob

## Usage

    k8s-safe-cronjob 'command'

Runs 'command' and always exits with 0, so that kubernetes will think the job has succeeded.

## Sentry

    k8s-safe-cronjob -dsn https://xxxx:yyyy@sentry.io/zzzzz 'command'

Will report errors to sentry, but then terminate with a zero return code.

