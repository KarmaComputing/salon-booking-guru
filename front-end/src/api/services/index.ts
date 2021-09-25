import { useAccountService } from './accountService';
import { useAuthenticateService } from './authenticateService';
import { useAvailabilityService } from './availabilityService';
import { useProductCategoryService } from './productCategoryService';
import { useProductService } from './productService';
import { useQualificationService } from './qualificationService';
import { useRoleService } from './roleService';

export const useService = () => {
    return {
        ...useAccountService(),
        ...useAuthenticateService(),
        ...useAvailabilityService(),
        ...useProductCategoryService(),
        ...useProductService(),
        ...useQualificationService(),
        ...useRoleService(),
    };
};
