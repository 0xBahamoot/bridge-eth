package bridge

// LOCAL - Beacon committees
var localBridgeCommitteePubKeys = []string{}

// var localBeaconCommitteePubKeys = []string{
//     "121VhftSAygpEJZ6i9jGkEKLMQTKTiiHzeUfeuhpQCcLZtys8FazpWwytpHebkAwgCxvqgUUF13fcSMtp5dgV1YkbRMj3z42TW2EebzAaiGg2DkGPodckN2UsbqhVDibpMgJUHVkLXardemfLdgUqWGtymdxaaRyPM38BAZcLpo2pAjxKv5vG5Uh9zHMkn7ZHtdNHmBmhG8B46UeiGBXYTwhyMe9KGS83jCMPAoUwHhTEXj5qQh6586dHjVxwEkRzp7SKn9iG1FFWdJ97xEkP2ezAapNQ46quVrMggcHFvoZofs1xdd4o5vAmPKnPTZtGTKunFiTWGnpSG9L6r5QpcmapqvRrK5SiuFhNM5DqgzUeHBb7fTfoiWd2N29jkbTGSq8CPUSjx3zdLR9sZguvPdnAA8g25cFPGSZt8aEnFJoPRzM",
//     "121VhftSAygpEJZ6i9jGkEqPGAXcmKffwMbzpwxnEfzJxen4oZKPukWAUBbqvV5xPnowZ2eQmAj2mEebG2oexebQPh1MPFC6vEZAk6i7AiRPrZmfaRrRVrBp4WXnVJmL3xK4wzTfkR2rZkhUmSZm112TTyhDNkDQSaBGJkexrPbryqUygazCA2eyo6LnK5qs7jz2RhhsWqUTQ3sQJUuFcYdf2pSnYwhqZqphDCSRizDHeysaua5L7LwS8fY7KZHhPgTuFjvUWWnWSRTmV8u1dTY5kcmMdDZsPiyN9WfqjgVoTFNALjFG8U4GMvzV3kKwVVjuPMsM2XqyPDVpdNQUgLnv2bJS8Tr22A9NgF1FQfWyAny1DYyY3N5H3tfCggsybzZXzrbYPPgokvEynac91y8hPkRdgKW1e7FHzuBnEisPuKzy",
//     "121VhftSAygpEJZ6i9jGkGLcYhJBeaJTGY5aFjqQA2WwyxU69Utrviuy9AJ3ATkeEyigVGScQUZw22cD1HeFKiyASYAs82WEamujt3nefYA9FPhURBpRTn6jDmGKUdb4QNbs7HVCJkRRaL9aktg1yaQaZE8TJFg2UeE9tBqUdmvD8fy36aDCYM5W86jaTVCXeEJQWPxUunP2EEL3e283PJ8zqPeBkpoFvkvhB28Hk3oRDeCCTC7QhbaV18ayKeToYqAxoUMBBihanfA33ixeX1daeKpajLCgDZ6jrfphwdYwQbf7dMcZ2NVvQ1a5JUCTJUZypwgKRt8tnTAKCowt2L1KNGP4NJJZm61cfHAGbKRyG9QxCJgK2SdMKsKPVefZSc9LbVaB7VeBby5LHxvMoCD7bN7g1HYRp4BX9n1fZJUeEkVa",
//     "121VhftSAygpEJZ6i9jGkDjJj7e2cfgQvrLsPsmLhGMmGD9U9Knffa1MZAw79EijnpueVfTStN2VYt5jRqEr2DTjVqzUinwHVKWH4Tg4szHUntiBdWeqzNC4E8iiwC9Y2KtcRr3hBkpfqvyuBvchigatrigRvFVWu8H2RQqjvopLL51DQ4LFD87L9Zgj9HhasMeyr6f37yirs47JgtGs4BM7EhhpM5zD3TCsFabPphtwDKnfuLMaGzoAw5fM8zEXvdLMuohk96oayjdYothncdtZom17DxB1Mmw535eEjxBwz9ELoZRKk3LYiheSd4xGN9QsxrT2WnZCTd8B5QktARte5S91QYvRMixKC8UEuovQhXt8jMZNkq7CmMeXoybfYdmNaAHuqbY1QeUT2AgaqPho4ay3z5eeKRhnB28H18RGWQ1L",
// }
var localBeaconCommitteePubKeys = []string{
	"121VhftSAygpEJZ6i9jGkD6teSvpg4YrDXnRasfFf7EnJfVdphCNYKQs9mZrsCgmhPjbwy3izcC8GYxU4A8gK2FtNAETCWqKeTpk9aYetJPcuaiVK6N9nMqPApx4Ja11HSyedE8QXgfptKXW8BdCEfDAGqmCFutrHWbmzSYEZ9wvysASu2GStHcHwY8PxansqdXWRHZmt8s4ocu2rcco7XmKFVfm8LtyA2KfBTwP7tVohCLYvXk7g9sfR4DYUzEGyWqe4iZg1zUXx75CkuQp1yvN4uRBPhVVWgg6JK1d76ffGbFsH1E4MfYenU4WJR9QJzW65ZWvjWG8mn4x7h37KGkCS47yu3x5uiuitKYg3su286VSy1vFqryhAF68nBs3xEgHrnx7rzH88Wq9Nba3TuP7Pmyme8u8dLJwPZTLSN391T9p",
	"121VhftSAygpEJZ6i9jGkQpWMhSRBSwuGA3dQX1R1444WAvvyeMSNENmWpAoR822buXaNChbnrBSHNeVoeaduTEBRbETWZLo2QPZxzDgNW8SKbpmuWyxZL97RK3bwJTy46orexKeiqyzovrwtenUN7HpKg9smaLEBWpvkRXpr6VZTtSS5My9VKWd4Sg8y9TPd84LWFsec1CG6tLugb2S2HEmJgoLyMMwPysjGCP7Hdjde7Hx4u9P3o8D4uthEu6JCFWrKoWkqSWtxzJi7cvX1nX7YWNNEUvjZJ86XuzL3GxcZ5UsdrTpmZWodSTsj4qqpx1PBCMQ3kesGAeAppiiJefPkJKiRHVhx9GxpwnLGUH8WMExYrAKF2oYsdZ3HyBSNBbyJfZSY5g91uFYY13CDU8AS7pUhwKYNcDAoPUpbL3EL3kc",
	"121VhftSAygpEJZ6i9jGk5W4zb8p2LSRMPhpUdpMsckbe5Gtt92Y9gY91z6WuiCdsrMVV66ZuEd6BTcfrQXaf3WqyayGXz8o5bUfAFg6YbubigWotxK1Rjz6yR6GCyxvDrXwvS5fRuRvjh9bhr4i3ayEdQrnYSNBkrY5Ykuj87fZ5NsTrD4yFEuUt63BwbH8aYiywMp9fVsmWuYRMS8cdv7epEEzyHiGADMhTTuaKfAcNtrXWCzfnEbzkEM3xvymdmWG1PVwfCQgNQ1FvPUB31pkTAHdEWUN6MrjRRTwifxt7pdE1sGLt2vusr6CQc4T6ELnpRjTQLTchDLUTUwB2u8fB1x6hVREzMGPuATpxcWawmnMa4Z6gjrZhrfN18htNzq3XxfEV27uX7nHYgFTFEV5468pCdUF3pT3hiU3Lm2TGZCs",
	"121VhftSAygpEJZ6i9jGkD6e1TFeFEsFEv4tjxLA9eRUxaT5fLBUSEA5ZKQqXexd5qBXoWoLz1ZTnyuBT5J3VWR995NnVQ93xy5WUTRzBeQU8uKjA6XkrMVzoh3F4dVyrFtFMdJHeR8GNLEax5t5ZJDTWupXeubHgkDpmPc5j9t3FFEYmfNq7nZVF1nfyDHEK8YNAmJWjrqKS2YinK9vR3S2hBQ1qQ4SXGDT2LrVt5KmozycPDjDBDcwjtnbiYhDYGvAUxvy18KJmhaBZwbLa4y3ksTvvNsYcTy2B5pQuHRXLa28jGmAtWM9nKFMiTpX3tKUjKg1L774q9CeKy6TcY1kLY9uyJd2pwRuoSf4PengKiZsdVe2wk266z3uBPQBdcq4j2225H9VGBKGmf6nqMzeyVDJZ1AtvpxBNzFj9gEQ2m9n",
}

// TESTNET2 - Beacon committees
var testnet2BridgeCommitteePubKeys = []string{}
var testnet2BeaconCommitteePubKeys = []string{
	"121VhftSAygpEJZ6i9jGkEKLMQTKTiiHzeUfeuhpQCcLZtys8FazpWwytpHebkAwgCxvqgUUF13fcSMtp5dgV1YkbRMj3z42TW2EebzAaiGg2DkGPodcXDm4nu1GAbvAkfHEyfQ8nX3RNpQxHzpBCg5mvtr6CTQbjVRFnhUg78D6oBkkeHFkNDUT383FVkDmttu6VxSKq7WARLGuuF44BR5WDKRM58MJ9uK7kQ6WsX1ZSiiFakyANp7epjhzdP4oGyKtEvtVmwqs1XiCMDHMWNqAdLgsE1LRMsxkRykMCmpxuVaq9iAWCsR6iDEdWmZ3PZEzEVBQ57doj8eqbbZuwdukPhE1Mqmk4EmptZWP3kpxUCq5S4cSe9B5JVnVsrnb8hhHKDz3Md1qS8ZaT1gEyYDWAVPKLKiMpRsLmhF2kFmVnHJK",
	"121VhftSAygpEJZ6i9jGkEqPGAXcmKffwMbzpwxnEfzJxen4oZKPukWAUBbqvV5xPnowZ2eQmAj2mEebG2oexebQPh1MPFC6vEZAk6i7AiRPrZmfaRrQnjRdsDCk4Md2QGLF2LQvDoFrKBTurLya5WgKsMLTjJDddjXnZg93bMyUzS3wDbtqDdvhFeWNsErN8bBBymhMMrW2s2W9HYxTcikrQ7A9HzWSdF95dzLkW8cp3qMFbB5z7Jm8WWZcX5f9pC3JBTeRa993hUTdp94o7dnPhAKhw39eAAYfwP5d9RLfss11Ft9f35pKpurGhpA8oEgrBStaruoumDoqABfwvURqFgE8zGhmDRTfyYcKha4ZPas9rZQPRodDGBBoYzv69yWxNx3VpD5Bo71cdQF5SJ2UteNoMjBtJqAeLFq13BAB6yKA",
	"121VhftSAygpEJZ6i9jGkGLcYhJBeaJTGY5aFjqQA2WwyxU69Utrviuy9AJ3ATkeEyigVGScQUZw22cD1HeFKiyASYAs82WEamujt3nefYA9FPhURBpRgPrsw3eCbNhHdtB64TAKNakonxGCBmq7AatSm9iQPKnf3xBToTTScR2SowJgY2yVpxe7odo6JWwbg8x74esix2Kp1rB8Y2Wm5nKnP3oFFSnkD3v9A9shE7w4xKdHBot7GmyD9KRoagCu5aojQnHxLNZYrEhUPSEmGZLqmnm7TBK2aUwJ9YfEhdked8H5Fa7CRaPQfuL2YA7PTeEERjttJP7B2uTCviUYCHwsfqhGzZy6s5FP2S7fd2DSKW5NWkZfEBvAUEn3wg36kvioa1pLKxVmz5BaXZt3PPcn4s2f2CdNup5p92EnV6jLBQTJ",
	"121VhftSAygpEJZ6i9jGkDjJj7e2cfgQvrLsPsmLhGMmGD9U9Knffa1MZAw79EijnpueVfTStN2VYt5jRqEr2DTjVqzUinwHVKWH4Tg4szHUntiBdWerKwQDjT3qFD7UtqhbZ7rfULf8fBvfYaqbaqecEZ1VKp3iMipvw4gj3tdhaX5s4wrTrEgd2TUEdCzkgmnDTPYhYF9ETuN29MzeStUKuL2ap7BH5e2yAdCdLoZuAihHdVB5aPB6bii54JoDW7P7uLHkzD3LoYySRmbxXBo31abDVM96wMxJ1gvfgg1Xswsxjrddwf3Vo55G1z88VY3zjHNfEy1ahkgTpm3mtjWP3EqM2SnKKijxcnz3jjNs2sMxkLcrrnbEK7uAuafMAEZuxi9jY8B8yedpya8hJQp7mx448wiuBXi3o9rcPfVL4UZd",
}

// MAINNET - Beacon committees
var mainnetBridgeCommitteePubKeys = []string{}
var mainnetBeaconCommitteePubKeys = []string{
	"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",
}
