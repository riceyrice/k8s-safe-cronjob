# k8s-safe-cronjob

## Usage

    k8s-safe-cronjob 'command'

Runs 'command' and always exits with 0, so that kubernetes will think the job has succeeded.

## Sentry

    k8s-safe-cronjob -dsn https://xxxx:yyyy@sentry.io/zzzzz 'command'

Will report errors to sentry, but then terminate with a zero return code.

## MaxFail

    k8s-safe-cronjob -fN 'command'

Runs 'command' and, if it fails, checks how many times it has failed recently.
If it has failed more than N times then it will instead return a success code.
This allows a command to be retried a few times, but then terminate
permanently.

The current number of failures is stored as an annotation on the cronjob.
