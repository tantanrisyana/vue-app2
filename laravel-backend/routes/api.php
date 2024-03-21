<?php

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;
use App\Http\Controllers\SiswaController;

Route::prefix('v1')->group(function () {
    Route::resource('siswa', SiswaController::class)->except(['create', 'edit']);
});
