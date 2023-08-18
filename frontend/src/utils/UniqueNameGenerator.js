import { faker } from '@faker-js/faker';



class UniqueNameGenerator {
    constructor() {
        this.namesSet = new Set();
        this.maxAttempts = 10; // Number of times to try for a unique name.
    }

    generateUniqueName() {
        let attempts = 0;
        let name;

        while (attempts < this.maxAttempts) {
            name = faker.internet.userName();
            if (!this.namesSet.has(name)) {
                this.namesSet.add(name);
                return name;
            }
            attempts++;
        }

        // Handle the edge case where a unique name isn't generated after maxAttempts.
        // You can modify this to suit your needs.
        throw new Error("Unable to generate a unique name");
    }
}

const generator = new UniqueNameGenerator();
export default generator

