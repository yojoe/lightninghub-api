let config = {
  bitcoind: {
    rpc: 'http://__cookie__:af49ac8cfa19d6ef0a100de3d526ace81a85f7f545537557ad8d5b22ea21cc15@127.0.0.1:18332',
  },
  lightningd: {
    rpc: 'http://127.0.0.1:19737',
    path: '/rpc',
    authkey: "masterkeythatcandoeverything"
  },
  redis: {
    port: 6379,
    host: '127.0.0.1',
    family: 4,
    password: '',
    db: 0,
  },
  lnd: {
    url: '1.1.1.1:10009',
    password: '',
  },
};

if (process.env.CONFIG) {
  console.log('using config from env');
  config = JSON.parse(process.env.CONFIG);
}

module.exports = config;
