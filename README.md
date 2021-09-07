# The cure

Goal of the project: a simple webapp to inspire people to quit bad habits costing them money,
for instead showing them what they could get with that money.
Let's begin with an average price of a pack of tobacco, converted to XTZ.
Then fetching NFT of the day that the price of such a pack could buy.

## Dev

You'll need to have these tools installed:

- [modd](https://github.com/cortesi/modd)
- [entr](https://github.com/clibs/entr) (optional)

Using tmuxinator:

```
cp .tmuxinator.yml.dist .tmuxinator.yml
tmuxinator start dev
```

Or manually:

Back: `modd`
Front: `npm --prefix ./assets_src run watch`
