module.exports = {
    plugins: [
        {
            serve(context) {
                if (context.path === '/teammates') {
                    return {body: [
                        { name: "Asdf", position: 1, status: 0 },
                        { name: "Qwer", position: 2, status: 1 },
                        { name: "Yxcv", position: 3, status: 2 },
                    ], type: 'application/json'}
                }
                if (context.path === '/matches') {
                    return {body: [
                        { date: "2020-09-09T12:00:00Z", description: "Test 1" },
                        { date: "2020-09-10T12:00:00Z", description: "Test 2" },
                    ], type: 'application/json'}
                }
                if (context.path === '/votes') {
                    return {body: [
                        {"teammate":"1","match":"20200909","vote":0},
                        {"teammate":"2","match":"20200909","vote":0},
                        {"teammate":"3","match":"20200909","vote":1},
                        {"teammate":"1","match":"20200910","vote":2},
                        {"teammate":"3","match":"20200910","vote":1}
                    ], type: 'application/json'}
                }
            }
        }
    ]
}