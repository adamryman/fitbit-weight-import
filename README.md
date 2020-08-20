# Fitbit Import Weight

I wanted to import my weight data from other apps into the fitbit api.

This is a collection of notes and scripts that I used.

Feel free to use them for reference, they will not work out of the box.

## Create a fitbit app

Register an app here for personal use: [https://dev.fitbit.com/apps](https://dev.fitbit.com/apps)

The urls you provide don't matter, as we are going to use the fitbit oauth interactive tutorial to get the authorization token we need.

http://example.com for most of the urls and http://example.com for the callback / redirect uri should work just fine.

## Request authorization token

Navigate to [https://dev.fitbit.com/apps/oauthinteractivetutorial](https://dev.fitbit.com/apps/oauthinteractivetutorial)

Enter the information from the fitbit app you created.

Follow the guide, in step 3 you will be provided a OAuth 2.0 Access Token, this will give you the value of `FITBIT_ACCESS_TOKEN=<TOKEN>` which is used in this repository.

# fitbit_import_weight

Even without setting `FITBIT_ACCESS_TOKEN`, you can use `fitbit_import_weight.sh` as follows:

```
FITBIT_ACCESS_TOKE=<TOKEN> ./fitbit_import_weight <date in yyyy-mm-dd format> <weight in pounds>
```
