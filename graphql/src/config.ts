import * as path from 'path';
import * as fs from "fs";
const config = {
    swagger: {
        options: {
            basePath: '/api/',
            documentationPath: '/docs',
            jsonEditor: true,
            info: {
                title: 'API Documentation',
                version: 'v1.0.0',
                contact: {
                    name: 'John doe',
                    email: 'johndoe@johndoe.com',
                },
            },
            grouping: 'tags',
            sortEndpoints: 'ordered',
        },
    },
    status: {
        path: '/status',
        options: {
            title: 'API Monitor',
            routeConfig: {
                auth: false,
            },
        },
    },
    datastore: {
        nedb: {
            users: {
                // inMemoryOnly: true,
                filename: path.join(process.cwd(), `storage/nedb/${process.env.NODE_ENV}/users.db`),
            },
            articles: {
                // inMemoryOnly: true,
                filename: path.join(process.cwd(), `storage/nedb/${process.env.NODE_ENV}/articles.db`),
            },
        },
    },
    grpc: {
        certs: {
            root: fs.readFileSync('/run/secrets/ca_root'),
            cert: fs.readFileSync('/run/secrets/client_crt'),
            key: fs.readFileSync('/run/secrets/client_key'),
        },
        services: {
            todos: {
                address: 'grpc-api:8081',
            }
        }
    }
};

export default config;
