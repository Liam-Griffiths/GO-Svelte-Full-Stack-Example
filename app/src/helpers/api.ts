import axios from 'axios';
import { toast } from '@zerodevx/svelte-toast'

const API_URL: string = "http://localhost:3000";

export type Product = {
    ProductID:        number;
    ProductName:      string;
    ProductOwnerName: string;
    Developers:       string[];
    ScrumMasterName:  string;
    StartDate:        Date;
    Methodology:      string;
}

export async function Health() {
    try {
        // 👇️ const data: CreateUserResponse
        const { data, status } = await axios.get(
            API_URL + '/api',
            {
                headers: {
                    'Content-Type': 'application/json',
                    Accept: 'application/json',
                },
            },
        );
;
        toast.push("API Online")
        return data;
    } catch (error) {
        if (axios.isAxiosError(error)) {
            console.log('error message: ', error.message);
            toast.push(error.message);
            return error.message;
        } else {
            console.log('unexpected error: ', error);
            return 'An unexpected error occurred';
        }
    }
}

export async function GetProducts(): Promise<Product[]> {
    try {
        // 👇️ const data: CreateUserResponse
        const { data, status } = await axios.get(
            API_URL + '/api/products',
            {
                headers: {
                    'Content-Type': 'application/json',
                    Accept: 'application/json',
                },
            },
        );
        ;
        toast.push("Products Fetched")
        return data;
    } catch (error) {
        if (axios.isAxiosError(error)) {
            console.log('error message: ', error.message);
            toast.push(error.message);
            return error.message;
        } else {
            console.log('unexpected error: ', error);
            return 'An unexpected error occurred';
        }
    }
}

export async function CreateProduct(newProduct: Product) {
    try {

        const { data, status } = await axios.post<Product>(
            API_URL + '/api/products',
            newProduct,
            {
                headers: {
                    'Content-Type': 'application/json',
                    Accept: 'application/json',
                },
            },
        );

        toast.push("Product successfully created")
        return data;
    } catch (error) {
        if (axios.isAxiosError(error)) {
            console.log('error message: ', error.message);
            toast.push(error.message)
            return error.message;
        } else {
            console.log('unexpected error: ', error);
            return 'An unexpected error occurred';
        }
    }
}

export async function UpdateProduct(newProduct: Product) {
    try {

        const { data, status } = await axios.put(
            API_URL + '/api/products/' + newProduct.ProductID,
            newProduct,
            {
                headers: {
                    'Content-Type': 'application/json',
                    Accept: 'application/json',
                },
            },
        );

        toast.push("Product successfully updated")
        return data;
    } catch (error) {
        if (axios.isAxiosError(error)) {
            console.log('error message: ', error.message);
            toast.push(error.message)
            return error.message;
        } else {
            console.log('unexpected error: ', error);
            return 'An unexpected error occurred';
        }
    }
}

export async function DeleteProduct(newProduct: Product) {
    try {

        const { data, status } = await axios.delete(
            API_URL + '/api/products/' + newProduct.ProductID,
            {
                headers: {
                    'Content-Type': 'application/json',
                    Accept: 'application/json',
                },
            },
        );

        toast.push("Product successfully deleted")
        return data;
    } catch (error) {
        if (axios.isAxiosError(error)) {
            console.log('error message: ', error.message);
            toast.push(error.message)
            return error.message;
        } else {
            console.log('unexpected error: ', error);
            return 'An unexpected error occurred';
        }
    }
}

