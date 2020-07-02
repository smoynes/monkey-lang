let map = fn(arr, f) {
    let iter = fn(arr, acc) {
      if (len(arr) == 0) {
        acc
      } else {
        let head = first(arr)
        let tail = rest(arr)
        iter(tail, push(acc, f(head)))
      }
    }

    iter(arr, [])
}

let reduce = fn(arr, initial, f) {
    let iter = fn(arr, result) {
      if (len(arr) == 0) {
        result
      } else {
        iter(rest(arr), f(result, first(arr)))
      }
  }

  iter(arr, initial)
}

let sum = fn(arr) {
    reduce(arr, 0, fn(acc, el) { acc + el })
}

let array = [0, 1, 2, 3]
puts(map(array, fn(x) { x * 2 }))
puts(reduce(array, 0, fn(acc, el) { acc + 1} ))
puts(sum(array))

let n = 0
while (n < 5) {
      n = n + 1;
}
