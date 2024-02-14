const data = require('./account.user_accounts.json')
const fs = require('fs')

const res = data.map(user => ({
    actor: {
        uuid: user.user_uuid,
        name: user.user_name,
        type:"user"
    },
    telegram: [],
    mail: [],
    sms: [],
    updatedAt: new Date().toISOString(),
}))

fs.writeFileSync('./output.user.json', JSON.stringify(res, null, 2))