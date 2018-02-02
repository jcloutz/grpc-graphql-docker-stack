import 'winston-daily-rotate-file';
import { interfaces } from 'inversify';
import { Logger, TransportInstance, transports } from 'winston';
import Config from '../config';
import { IConfig, ILogger } from '../interfaces';
import Types from '../ioc/types';
import * as todoService from "../protobuf/todo_grpc_pb";
import * as grpc from 'grpc';
import {ITodoClient} from "../protobuf/todo_grpc_pb";

export default (c: interfaces.Container): void => {
    c.bind<interfaces.Container>(Types.Services.Container).toConstantValue(c);

    c.bind<IConfig>(Types.Services.Config).toConstantValue(Config);

    c.bind<TransportInstance>('LoggerTransport').toConstantValue(
        new transports.DailyRotateFile({
            level: process.env.LOG_LEVEL,
            datePattern: 'dd-MM-yyyy.',
            dirname: './logs',
            filename: './log',
            prepend: true,
        }),
    );

    c.bind<TransportInstance>('LoggerTransport').toConstantValue(
        new transports.Console({
            colorize: true,
            prettyPrint: true,
            level: process.env.NODE_ENV === 'test' ? 'warn' : 'info',
        }),
    );

    c.bind<ILogger>(Types.Services.Logger).toDynamicValue((context: interfaces.Context) => {
        return new Logger({
            transports: context.container.getAll('LoggerTransport'),
        });
    });

    c.bind<ITodoClient>(Types.Services.Todos).toDynamicValue((context: interfaces.Context) => {
        const config = context.container.get<IConfig>(Types.Services.Config);
        return new todoService.TodoClient('localhost:8081', grpc.credentials.createSsl(
            config.grpc.certs.root,
            config.grpc.certs.key,
            config.grpc.certs.cert
        ));
    }).inSingletonScope();
};
