import * as Joi from 'joi';

export default {
    create: {
        payload: {
            name: Joi.string().required(),
            completed: Joi.boolean().required(),
        },
    },
    updateById: {
        params: {
            id: Joi.string().required(),
        },
        payload: {
            name: Joi.string().required(),
            completed: Joi.boolean().required()
        },
    },
    getById: {
        params: {
            id: Joi.string().required(),
        },
    },
    deleteById: {
        params: {
            id: Joi.string().required(),
        },
    },
};
