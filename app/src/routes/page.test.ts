import {describe, test, expect} from 'vitest';
import { render, screen } from '@testing-library/svelte';

import Page from "./+page.svelte";

describe("Page Component", () => {

    test("should render the component with the app title", () => {
        render(Page);
        expect(screen.getByText('IMB App Manager')).toBeTruthy()
    });

})