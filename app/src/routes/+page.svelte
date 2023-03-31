<script lang="ts">
    import { writable } from 'svelte/store';
    import Modal, {bind} from 'svelte-simple-modal';
    import { onMount } from 'svelte';

    import Popup from './Popup.svelte';
    import type {Product} from "../helpers/api";
    import {GetProducts, Health} from "../helpers/api";

    let products: Product[] = [];
    let productsCopy: Product[] = [];
    let isLoading = true;
    let term = "";
    let sortOrder = 1;

    const modal = writable(null);
    const showModal = (product: Product, isNew = false) => modal.set(bind(Popup, { product: product, isNew: isNew, fetchFunc: FetchProducts }));

    onMount(async () => {
        await Health()
        products = await GetProducts();
        productsCopy = [...products];
        isLoading = false;
    });

    async function FetchProducts() {
        products = await GetProducts();
    }

    const truncate = (str, length) => {
        if (length == null) {
            length = 100;
        }

        let ending = '...';

        if (str.length > length) {
            return str.substring(0, length - ending.length) + ending;
        } else {
            return str;
        }
    };;

    const dynamicSort = property => {
        console.log(JSON.stringify(products))
        if(sortOrder === -1){
            sortOrder = 1
        }else{
            sortOrder = -1
        }

        return function (a,b) {
            let result = (a[property] < b[property]) ? -1 : (a[property] > b[property]) ? 1 : 0;
            return result * sortOrder;
        }
    };

    const doSearch = () => {
        products = productsCopy.filter(item => {
                if(Object.values(item).toString().toLowerCase().includes(term.toLowerCase())){
                    return true;
                }
                return false;
        });
    };
</script>

<Modal show={$modal}>
</Modal>

<section class="bg-white dark:bg-gray-900">
    <div class="py-2 px-4">
        <div class="mx-auto max-w-screen-sm text-center lg:mb-4 mb-2 mt-2">
            <h2 class="text-2xl lg:text-3xl tracking-tight font-extrabold text-gray-900 dark:text-white">IMB Management Portal</h2>
        </div>
    </div>
</section>

<section class="bg-gray-50 dark:bg-gray-900 flex items-center">
    <div class="max-w-screen-xl px-4 mx-auto w-full">

        <div class="relative bg-white shadow-md dark:bg-gray-800 sm:rounded-lg mt-5">
            <div class="flex flex-col items-center justify-between p-4 space-y-3 md:flex-row md:space-y-0 md:space-x-4">
                          <span class="text-sm font-normal text-gray-500 dark:text-gray-400">Total:<span
                                  class="font-semibold text-gray-900 dark:text-white"> {products.length ?? 0}</span></span>
                <div class="w-full">
                    <form class="flex items-center">
                        <label for="simple-search" class="sr-only">Search</label>
                        <div class="relative w-full">
                            <div class="absolute inset-y-0 left-0 flex items-center pl-3 pointer-events-none">
                                <svg aria-hidden="true" class="w-5 h-5 text-gray-500 dark:text-gray-400" fill="currentColor" viewbox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
                                    <path fill-rule="evenodd" d="M8 4a4 4 0 100 8 4 4 0 000-8zM2 8a6 6 0 1110.89 3.476l4.817 4.817a1 1 0 01-1.414 1.414l-4.816-4.816A6 6 0 012 8z" clip-rule="evenodd" />
                                </svg>
                            </div>
                            <input type="text" on:input={(e) => doSearch()} bind:value={term} id="simple-search" class="block w-full p-2 pl-10 text-sm text-gray-900 border border-gray-300 rounded-lg bg-gray-50 focus:ring-primary-500 focus:border-primary-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500" placeholder="Search" required="">
                        </div>
                    </form>
                </div>
                <div on:click={() => showModal({}, true)} class="flex flex-col items-stretch justify-end flex-shrink-0 w-full space-y-2 md:w-auto md:flex-row md:space-y-0 md:items-center md:space-x-3">
                    <button type="button" class="flex items-center justify-center px-4 py-2 text-sm font-medium text-white rounded-lg bg-primary-700 hover:bg-primary-800 focus:ring-4 focus:ring-primary-300 dark:bg-primary-600 dark:hover:bg-primary-700 focus:outline-none dark:focus:ring-primary-800">
                        <svg class="h-3.5 w-3.5 mr-2" fill="currentColor" viewbox="0 0 20 20" xmlns="http://www.w3.org/2000/svg" aria-hidden="true">
                            <path clip-rule="evenodd" fill-rule="evenodd" d="M10 3a1 1 0 011 1v5h5a1 1 0 110 2h-5v5a1 1 0 11-2 0v-5H4a1 1 0 110-2h5V4a1 1 0 011-1z" />
                        </svg>
                        Add product
                    </button>

                </div>
            </div>
        </div>
        <div class="relative bg-white shadow-md dark:bg-gray-800 sm:rounded-lg mt-5">
            <div class="flex flex-col items-center justify-between p-4 space-y-3 md:flex-row md:space-y-0 md:space-x-4">
                <div class="w-full">
                    <table class="w-full text-sm text-left text-gray-500 dark:text-gray-400">
            <thead class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
            <tr>
                <th scope="col" class="px-6 py-3">
                    <a on:click={() => {products = products.sort(dynamicSort("ProductID"))}}>Product ID</a>
                </th>

                <th scope="col" class="px-6 py-3">
                    <a on:click={() => {products = products.sort(dynamicSort("ProductName"))}}>Product Name</a>
                </th>

                <th scope="col" class="px-6 py-3">
                    <a on:click={() => {products = products.sort(dynamicSort("ProductOwnerName"))}}>Product Owner</a>
                </th>

                <th scope="col" class="px-6 py-3">
                    Developers
                </th>

                <th scope="col" class="px-6 py-3">
                    <a on:click={() => {products = products.sort(dynamicSort("ScrumMasterName"))}}>Scrum Master</a>
                </th>

                <th scope="col" class="px-6 py-3">
                    <a on:click={() => {products = products.sort(dynamicSort("StartDate"))}}>Start Date</a>
                </th>

                <th scope="col" class="px-6 py-3">
                    <a on:click={() => {products = products.sort(dynamicSort("Methodology"))}}>Methodology</a>
                </th>

                <th scope="col" class="px-6 py-3">
                    <span class="sr-only">Edit</span>
                </th>
            </tr>
            </thead>
            <tbody>
            {#if isLoading}
                <h1>Loading...</h1>
            {:else}
                {#if products.length}

                    {#each products as product}
                        <tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
                            <th scope="row" class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap dark:text-white">
                                {product.ProductID}
                            </th>
                            <th scope="row" class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap dark:text-white">
                                {product.ProductName}
                            </th>
                            <td class="px-6 py-4">
                                {product.ProductOwnerName}
                            </td>
                            <td class="px-6 py-4">
                                {product.Developers?.length ? product.Developers.join(", ") : '' }
                            </td>
                            <td class="px-6 py-4">
                                {product.ScrumMasterName}
                            </td>
                            <td class="px-6 py-4">
                                {new Date(product.StartDate).toString()}
                            </td>
                            <td class="px-6 py-4">
                                {product.Methodology}
                            </td>
                            <td class="px-6 py-4 text-right">
                                <a href="#" on:click={showModal(product)} class="font-medium text-blue-600 dark:text-blue-500 hover:underline">Edit</a>
                            </td>
                        </tr>
                    {/each}

                {/if}
            {/if}

            </tbody>
        </table>
                </div>
            </div>
        </div>
    </div>
</section>