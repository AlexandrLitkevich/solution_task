

function swapValues  () {
    const args = Array.prototype.slice.call(arguments);
    const arr = args[0]
    const temp = arr[0];
    arr[0] = arr[1];
    arr[1] = temp;
    args[0] = arr
}



swapValues([1,2])