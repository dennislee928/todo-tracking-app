import 'dart:async';
import 'dart:io';

import 'package:flutter/foundation.dart';
import 'package:device_shared/services/api_service.dart';
import 'package:in_app_purchase/in_app_purchase.dart';

/// Product ID for remove ads - must match App Store Connect / Google Play Console
const kRemoveAdsProductId = 'remove_ads';

class SubscriptionService {
  final ApiService _api = ApiService();
  final InAppPurchase _iap = InAppPurchase.instance;

  StreamSubscription<List<PurchaseDetails>>? _subscription;

  /// Called when purchase is verified - consumer should refetch user state
  VoidCallback? onPurchaseVerified;

  Future<void> init() async {
    if (!await _iap.isAvailable()) return;

    _subscription = _iap.purchaseStream.listen(
      _onPurchaseUpdate,
      onDone: () => _subscription?.cancel(),
      onError: (_) => _subscription?.cancel(),
    );
  }

  void dispose() {
    _subscription?.cancel();
  }

  Future<List<ProductDetails>> getProducts() async {
    if (!await _iap.isAvailable()) return [];
    final ids = {kRemoveAdsProductId};
    final response = await _iap.queryProductDetails(ids);
    if (response.notFoundIDs.isNotEmpty) return [];
    return response.productDetails;
  }

  Future<bool> purchaseRemoveAds() async {
    if (!await _iap.isAvailable()) return false;

    final products = await getProducts();
    if (products.isEmpty) return false;

    final productDetails = products.first;
    final purchaseParam = PurchaseParam(productDetails: productDetails);

    return _iap.buyNonConsumable(purchaseParam: purchaseParam);
  }

  Future<void> _onPurchaseUpdate(List<PurchaseDetails> purchases) async {
    for (final purchase in purchases) {
      if (purchase.status == PurchaseStatus.pending) continue;
      if (purchase.status == PurchaseStatus.error) continue;
      if (purchase.status == PurchaseStatus.canceled) continue;

      if (purchase.status == PurchaseStatus.purchased ||
          purchase.status == PurchaseStatus.restored) {
        await _verifyPurchase(purchase);
      }
    }
  }

  Future<void> _verifyPurchase(PurchaseDetails purchase) async {
    try {
      if (!Platform.isIOS && !Platform.isAndroid) return;
      if (Platform.isIOS) {
        final receipt = purchase.verificationData.serverVerificationData;
        if (receipt.isNotEmpty) {
          await _api.verifyApplePurchase(receipt);
          onPurchaseVerified?.call();
        }
      } else if (Platform.isAndroid) {
        final token = purchase.verificationData.serverVerificationData;
        if (token.isNotEmpty) {
          await _api.verifyGooglePurchase(token, purchase.productID);
          onPurchaseVerified?.call();
        }
      }
    } catch (_) {
      // Verification failed - don't complete
    }
  }
}
