/* globals gauge*/
"use strict";
const path = require('path');
const {
    openBrowser,
    write,
    closeBrowser,
    goto,
    press,
    screenshot,
    above,
    click,
    checkBox,
    listItem,
    toLeftOf,
    link,
    text,
    into,
    textBox,
    evaluate
} = require('taiko');
const assert = require("assert");
const headless = process.env.headless_chrome.toLowerCase() === 'true';

beforeSuite(async () => {
    await openBrowser({
        headless: headless
    })
});

afterSuite(async () => {
    await closeBrowser();
});

// Return a screenshot file name
gauge.customScreenshotWriter = async function () {
    const screenshotFilePath = path.join(process.env['gauge_screenshots_dir'],
        `screenshot-${process.hrtime.bigint()}.png`);

    await screenshot({
        path: screenshotFilePath
    });
    return path.basename(screenshotFilePath);
};

step("Open shopping-cart application", async () => {
    await goto("localhost:3000");
});

step("Add first product to basket", async () => {
    await click("Add to Basket");
});

step("Go to <item> page", async (item) => {
    await click(item);
});

step("See the product", async () => {

});

step("Write to <item1> area <item2>", async (item1, item2) => {
    await write(item2, into(textBox(item1)));
});

step("Create the product", async () => {
    await click("Create");
});

step("<item1> <item2> times countity of product", async (item1, item2) => {
    await write(item1, into(textBox("Quantity")));
    await click(item2);
});

step("Delete product", async () => {
    await click("Delete");
})