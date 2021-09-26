import { useAccountService } from './accountService';
import { useAuthenticateService } from './authenticateService';
import { useAvailabilityService } from './availabilityService';
import { useProductCategoryService } from './productCategoryService';
import { useProductService } from './productService';
import { useQualificationService } from './qualificationService';
import { useRoleService } from './roleService';

export const useService = () => {
    let services: any = {
        ...useAccountService(),
        ...useAuthenticateService(),
        ...useAvailabilityService(),
        ...useProductCategoryService(),
        ...useProductService(),
        ...useQualificationService(),
        ...useRoleService(),
    };

    const generateDataServices = (modelName: string, isSummary: boolean) => {
        return {
            getAll: services[`getAll${modelName}${isSummary ? 'Summary' : ''}`],
            get: services[`get${modelName}`],
            create: services[`create${modelName}`],
            update: services[`update${modelName}`],
            delete: services[`delete${modelName}`],
        };
    };

    services = { ...services, generateDataServices };

    return services as any;
};
