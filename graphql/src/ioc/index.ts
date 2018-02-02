import { Container } from 'inversify';
import controllers from '../ioc/controllers';
import datastores from '../ioc/datastore';
import factories from '../ioc/factories';
import services from '../ioc/services';
import plugins from '../ioc/plugins';
import repositories from '../ioc/repositories';
import resolvers from '../ioc/resolvers';
import router from '../ioc/router';

const container = new Container();

// Load Services first for general availability
services(container);

router(container);
controllers(container);
datastores(container);
factories(container);
plugins(container);
repositories(container);
resolvers(container);

export default container;
