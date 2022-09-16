const ETH = '0x0000000000000000000000000000000000000000';
const pETH = 'ffd8d42dc40a8d166ea4848baf8b5f6e9fe0e9c30d60062eb7d44a8df9e00854';

// test tokenList
const pTokens = {
  '0x0000000000000000000000000000000000000000': 'ffd8d42dc40a8d166ea4848baf8b5f6e9fe0e9c30d60062eb7d44a8df9e00854',
  '0x64B464037Ef0Aa3D1a95A5c04BC77e8667870E90': '0000000000000000000000000000000000000000000000000000000000000001',
  '0xBb46A9BFA0B7fD28D623150d0Ae070020D6Aab56': '0000000000000000000000000000000000000000000000000000000000000002',
  '0xcabda2B57fa84F02AE0aE6eA3634db4F030014c1': '0000000000000000000000000000000000000000000000000000000000000003',
};

const unify = {
  '0x0000000000000000000000000000000000000000': 'b366fa400c36e6bbcf24ac3e99c90406ddc64346ab0b7ba21e159b83d938812d',
  '0x64B464037Ef0Aa3D1a95A5c04BC77e8667870E90': '000000000000000000000000000000000000000000000000000000000000ff01',
  '0xBb46A9BFA0B7fD28D623150d0Ae070020D6Aab56': '000000000000000000000000000000000000000000000000000000000000ff02',
  '0xcabda2B57fa84F02AE0aE6eA3634db4F030014c1': '000000000000000000000000000000000000000000000000000000000000ff03',
};
const networkIDs = {
  ETH: 1,
  BSC: 2,
  PLG: 3,
  FTM: 4,
}

module.exports = {
  tokenAddresses: {
    ETH,
    pETH,
    pTokens,
    unify,
  },
}