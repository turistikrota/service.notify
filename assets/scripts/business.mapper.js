const data = require('./business.businesses.json')
const fs = require('fs')

const res = data.map(business => ({
    actor: {
        uuid: business._id.$oid,
        name: business.nick_name,
        type:"business"
    },
    telegram: [],
    mail: [],
    sms: [],
    updatedAt: new Date().toISOString(),
}))

fs.writeFileSync('./output.business.json', JSON.stringify(res, null, 2))