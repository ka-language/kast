#ifndef OMM_GOATV_HELPER_H_
#define OMM_GOATV_HELPER_H_

#ifdef __cplusplus
extern "C" {
#endif

static inline void* getidx(void** arr, int idx) {
    return arr[idx];
}

#ifdef __cplusplus
}
#endif

#endif