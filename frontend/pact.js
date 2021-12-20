const path = require("path")
const { Publisher } = require("@pact-foundation/pact")

const opts = {
    pactFilesOrDirs: [path.resolve("pact/pacts/shoppingcartconsumer-shoppingcartprovider.json")],
    pactBroker: 'https://enesbek.pactflow.io',
    pactBrokerToken: 'BB6cjlL4fdMPxwxdQS6uLA',
    consumerVersion: '2.0.2'
};

new Publisher(opts)
  .publishPacts()
  .then(() => {
    console.log("Pact published!")
}).catch((error) => {
    console.log(error)
})